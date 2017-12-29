package status

import (
	"fmt"

	"k8s.io/client-go/pkg/api/v1"
)

const (
	// NodeUnreachablePodReason is the reason and message set on a pod
	// when its state cannot be confirmed as kubelet is unresponsive
	// on the node it is (was) running.
	// copy from k8s.io/kubernetes/pkg/util/node
	NodeUnreachablePodReason = "NodeLost"
	// PodInitializing means pod's initContainers are not finished
	PodInitializing v1.PodPhase = "Initializing"
	// PodError means that:
	//   - When pod is initializing, at least one init container is terminated without code 0.
	//   - When pod is terminating, at least one container is terminated without code 0.
	PodError v1.PodPhase = "Error"
)

// PodStatus represents the current status of a pod
type PodStatus struct {
	Name            string      `json:"name"`
	Namespace       string      `json:"namespace"`
	Ready           bool        `json:"ready"`
	RestartCount    int32       `json:"restartCount"`
	InitContainers  int32       `json:"initContainers"`
	ReadyContainers int32       `json:"readyContainers"`
	TotalContainers int32       `json:"totalContainers"`
	NodeName        string      `json:"nodeName"`
	Phase           v1.PodPhase `json:"phase"`
	Reason          string      `json:"reason"`
	Message         string      `json:"message,omitempty"`
}

// JudgePodStatus judges the current status of pod from Pod.Status
func JudgePodStatus(pod *v1.Pod) PodStatus {
	if pod == nil {
		return PodStatus{}
	}
	ready := false
	restarts := 0
	readyContainers := 0
	initContainers := len(pod.Spec.InitContainers)
	totalContainers := len(pod.Spec.Containers)
	phase := pod.Status.Phase
	reason := string(pod.Status.Phase)
	if pod.Status.Reason != "" {
		reason = pod.Status.Reason
	}
	message := ""

	if phase == v1.PodPending {
		// detect pending error
		for i := range pod.Status.Conditions {
			condition := pod.Status.Conditions[i]
			// unschedulable error
			if condition.Type == v1.PodScheduled &&
				condition.Status == v1.ConditionFalse &&
				condition.Reason == v1.PodReasonUnschedulable {
				phase = PodError
				reason = condition.Reason
				message = condition.Message
			}
		}
	}

	initializing := false
	for i := range pod.Status.InitContainerStatuses {
		container := pod.Status.InitContainerStatuses[i]
		restarts += int(container.RestartCount)
		switch {
		case container.State.Terminated != nil && container.State.Terminated.ExitCode == 0:
			// initialized success
			continue
		case container.State.Terminated != nil:
			// initialization is failed
			reason = fmt.Sprintf("Init:ExitCode:%d", container.State.Terminated.ExitCode)
			if container.State.Terminated.Signal != 0 {
				reason = fmt.Sprintf("Init:Signal:%d", container.State.Terminated.Signal)
			}
			if len(container.State.Terminated.Reason) > 0 {
				reason = "Init:" + container.State.Terminated.Reason
			}
			message = container.State.Terminated.Message
			phase = PodError
			initializing = true
		case container.State.Waiting != nil && len(container.State.Waiting.Reason) > 0 && container.State.Waiting.Reason != "PodInitializing":
			reason = "Init:" + container.State.Waiting.Reason
			message = container.State.Waiting.Message
			phase = PodInitializing
			initializing = true
		default:
			reason = fmt.Sprintf("Init:%d/%d", i, len(pod.Spec.InitContainers))
			message = string(PodInitializing)
			phase = PodInitializing
			initializing = true
		}
		break
	}

	if !initializing {
		for i := len(pod.Status.ContainerStatuses) - 1; i >= 0; i-- {
			container := pod.Status.ContainerStatuses[i]
			restarts += int(container.RestartCount)

			if container.State.Waiting != nil && container.State.Waiting.Reason != "" {
				reason = container.State.Waiting.Reason
				message = container.State.Waiting.Message
			} else if container.State.Terminated != nil {
				if container.State.Terminated.ExitCode != 0 {
					phase = PodError
				}
				reason = fmt.Sprintf("ExitCode:%d", container.State.Terminated.ExitCode)
				message = container.State.Terminated.Message

				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Signal:%d", container.State.Terminated.Signal)
				}
				if container.State.Terminated.Reason != "" {
					reason = container.State.Terminated.Reason
				}
			} else if container.Ready && container.State.Running != nil {
				readyContainers++
			}
		}
	}

	if readyContainers == totalContainers && readyContainers > 0 {
		ready = true
	}

	if pod.DeletionTimestamp != nil {
		ready = false
		if pod.Status.Reason == NodeUnreachablePodReason {
			phase = v1.PodUnknown
			reason = "Unknown"
		} else {
			reason = "Terminating"
		}
	}
	return PodStatus{
		Name:            pod.Name,
		Namespace:       pod.Namespace,
		Ready:           ready,
		RestartCount:    int32(restarts),
		ReadyContainers: int32(readyContainers),
		InitContainers:  int32(initContainers),
		TotalContainers: int32(totalContainers),
		NodeName:        pod.Spec.NodeName,
		Phase:           phase,
		Reason:          reason,
		Message:         message,
	}
}
