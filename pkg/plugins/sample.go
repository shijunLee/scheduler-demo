package plugins

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework"
	frameworkruntime "k8s.io/kubernetes/pkg/scheduler/framework/runtime"
)

// 插件名称
const Name = "sample-plugin"

type Args struct {
	FavoriteColor  string `json:"favorite_color,omitempty"`
	FavoriteNumber int    `json:"favorite_number,omitempty"`
	ThanksTo       string `json:"thanks_to,omitempty"`
}

var _ framework.PreFilterPlugin = &Sample{}

type Sample struct {
	args   *Args
	handle framework.Handle
}

func (s *Sample) Name() string {
	return Name
}
func (s *Sample) PreFilterExtensions() framework.PreFilterExtensions {
	return s
}

func (s *Sample) AddPod(ctx context.Context, state *framework.CycleState, podToSchedule *v1.Pod,
	podInfoToAdd *framework.PodInfo, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("addPod pod: %v", podToSchedule.Name)
	return framework.NewStatus(framework.Success, "")
}

// RemovePod is called by the framework while trying to evaluate the impact
// of removing podToRemove from the node while scheduling podToSchedule.
func (s *Sample) RemovePod(ctx context.Context, state *framework.CycleState, podToSchedule *v1.Pod,
	podInfoToRemove *framework.PodInfo, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("removePod pod: %v", podToSchedule.Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) PreFilter(ctx context.Context, state *framework.CycleState, p *v1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", p.Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, nodeInfo.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *Sample) PreBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("start pre bind pod %s/%s: %+v", p.Name, p.Namespace, nodeName)
	if nodeInfo, err := s.handle.SnapshotSharedLister().NodeInfos().Get(nodeName); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("prebind get node info error: %+v", nodeName))
	} else {
		klog.V(3).Infof("prebind node info: %+v", nodeInfo.Node())
		return framework.NewStatus(framework.Success, "")
	}
}

//func DecodeInto(obj runtime.Object, into interface{}) error {
//type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(obj runtime.Object, f framework.Handle) (framework.Plugin, error) {

	args := &Args{}
	if err := frameworkruntime.DecodeInto(obj, args); err != nil {
		klog.Error("merge config error", err)
		return nil, err
	}

	klog.V(3).Infof("get plugin config args: %+v", args)
	return &Sample{
		args:   args,
		handle: f,
	}, nil
}
