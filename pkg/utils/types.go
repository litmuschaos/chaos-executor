package utils

import (
	"flag"

	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	clientV1alpha1 "github.com/litmuschaos/chaos-operator/pkg/client/clientset/versioned"
	volume "github.com/litmuschaos/elves/kubernetes/volume/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"

	"github.com/litmuschaos/chaos-runner/pkg/utils/k8s"
	"github.com/litmuschaos/chaos-runner/pkg/utils/litmus"
)

// EngineDetails struct is for collecting all the engine-related details
type EngineDetails struct {
	Name             string
	Experiments      []string
	AppLabel         string
	SvcAccount       string
	AppKind          string
	ClientUUID       string
	AuxiliaryAppInfo string
	UID              string
	EngineNamespace  string
}

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	Name               string
	Env                map[string]string
	ExpLabels          map[string]string
	ExpImage           string
	ExpImagePullPolicy corev1.PullPolicy
	ExpArgs            []string
	JobName            string
	Namespace          string
	ConfigMaps         []v1alpha1.ConfigMap
	Secrets            []v1alpha1.Secret
	VolumeOpts         VolumeOpts
	SvcAccount         string
	Annotations        map[string]string
	SecurityContext    v1alpha1.SecurityContext
	HostPID            bool
}

//VolumeOpts is a strcuture for all volume related operations
type VolumeOpts struct {
	VolumeMounts   []corev1.VolumeMount
	VolumeBuilders []*volume.Builder
}

// ClientSets is a collection of clientSets needed
type ClientSets struct {
	KubeClient   *kubernetes.Clientset
	LitmusClient *clientV1alpha1.Clientset
}

// Recorder is collection of resources needed to record events for chaos-runner
type Recorder struct {
	EventRecorder record.EventRecorder
	EventResource runtime.Object
}

var (
	// DefaultExpImagePullPolicy contains the defaults value (Always) of imagePullPolicy for exp container
	DefaultExpImagePullPolicy corev1.PullPolicy = "Always"
)

const (
	ExperimentDependencyCheckReason  string = "ExperimentDependencyCheck"
	ExperimentJobCreateReason        string = "ExperimentJobCreate"
	ExperimentJobCleanUpReason       string = "ExperimentJobCleanUp"
	ExperimentSkippedReason          string = "ExperimentSkipped"
	ExperimentEnvParseErrorReason    string = "EnvParseError"
	ExperimentNotFoundErrorReason    string = "ExperimentNotFound"
	ExperimentJobCreationErrorReason string = "JobCreationError"
	ExperimentJobWatchErrorReason    string = "JobWatchNotPermitted"
)

// GenerateClientSetFromKubeConfig will generation both ClientSets (k8s, and Litmus)
func (clientSets *ClientSets) GenerateClientSetFromKubeConfig() error {
	config, err := getKubeConfig()
	if err != nil {
		return err
	}
	k8sClientSet, err := k8s.GenerateK8sClientSet(config)
	if err != nil {
		return err
	}
	litmusClientSet, err := litmus.GenerateLitmusClientSet(config)
	if err != nil {
		return err
	}
	clientSets.KubeClient = k8sClientSet
	clientSets.LitmusClient = litmusClientSet

	return nil
}

// Generate

// getKubeConfig setup the config for access cluster resource
func getKubeConfig() (*rest.Config, error) {
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	// Use in-cluster config if kubeconfig path is specified
	if *kubeconfig == "" {
		config, err := rest.InClusterConfig()
		if err != nil {
			return config, err
		}
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return config, err
	}
	return config, err
}
