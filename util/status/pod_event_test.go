package status

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/core/v1"
)

func TestJudgePodStatus(t *testing.T) {
	tests := []struct {
		name   string
		pod    *v1.Pod
		events []*v1.Event
		want   PodStatus
	}{
		{
			"running",
			&v1.Pod{
				Spec: v1.PodSpec{
					Containers: []v1.Container{{Name: "ready"}},
				},
				Status: v1.PodStatus{
					Phase: v1.PodRunning,
					ContainerStatuses: []v1.ContainerStatus{
						{Ready: true, State: v1.ContainerState{Running: &v1.ContainerStateRunning{}}},
					},
				},
			},
			nil,
			PodStatus{Ready: true, ReadyContainers: 1, TotalContainers: 1, State: PodNormal, Phase: PodRunning, Reason: string(v1.PodRunning)},
		},
		{
			"Readiness probe failed",
			&v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					UID:       "poduid",
					Name:      "podName",
					Namespace: "default",
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{Name: "ready"}},
				},
				Status: v1.PodStatus{
					Phase: v1.PodRunning,
					ContainerStatuses: []v1.ContainerStatus{
						{Ready: true, State: v1.ContainerState{Running: &v1.ContainerStateRunning{}}},
					},
				},
			},
			[]*v1.Event{
				&v1.Event{
					InvolvedObject: v1.ObjectReference{
						Kind:      "Pod",
						UID:       "poduid",
						Name:      "podName",
						Namespace: "default",
					},
					LastTimestamp: metav1.Now(),
					Type:          v1.EventTypeWarning,
					Reason:        EventUnhealthy,
					Message:       "Readiness probe failed: Get http://192.168.68.97:8080/: dial tcp 192.168.68.97:8080: getsockopt: connection refused",
				},
			},
			PodStatus{Ready: false, ReadyContainers: 1, TotalContainers: 1, State: PodAbnormal, Phase: PodError, Reason: EventUnhealthy, Message: "Readiness probe failed: Get http://192.168.68.97:8080/: dial tcp 192.168.68.97:8080: getsockopt: connection refused"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JudgePodStatus(tt.pod, tt.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JudgePodStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
