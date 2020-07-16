package utils

import (
	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

//NOTE: The hostFileVolumeUtils doesn't contain the function to derive hostFileVols from chaosengine
//and thereby, the corresponding ones to override chaosengine values over experiment. 
//This is because, the hostfiles mounted into exp are often for a very specific purpose, such as,
//socket file mounts etc., and are often have fixed paths, i.e., similar to securityContext/hostPID
//and other such mandatory attributes 


//PatchHostFileVolumes patches hostFileVolume in experimentDetails struct.
func (expDetails *ExperimentDetails) PatchHostFileVolumes(clients ClientSets, engineDetails EngineDetails) error {
	err := expDetails.SetHostFileVolumes(clients, engineDetails)
	if err != nil {
		return err
	}

	klog.V(0).Infof("Validating HostFileVolumes details specified in the ChaosExperiment")
	err = expDetails.ValidateHostFileVolumes(clients)
	if err != nil {
		return err
	}
	return nil
}

// SetHostFileVolumes sets the value of hostFileVolumes in Experiment Structure
func (expDetails *ExperimentDetails) SetHostFileVolumes(clients ClientSets, engineDetails EngineDetails) error {

	experimentHostFileVolumes, err := getExperimentHostFileVolumes(clients, expDetails)
	if err != nil {
		return err
	}

	expDetails.HostFileVolumes = experimentHostFileVolumes

	return nil
}

// ValidateHostFileVolumes validates the hostFileVolume definition in experiment CR spec
func (expDetails *ExperimentDetails) ValidateHostFileVolumes(clients ClientSets) error {

    for _, v := range expDetails.HostFileVolumes {
		if v.Name == "" || v.MountPath == "" || v.NodePath == "" {
			return errors.New("Incomplete Information in HostFileVolume, will skip execution")
		}
		klog.V(0).Infof("Successfully Validated HostFileVolume: %v", v.Name)
   }
	return nil
}

// getExperimentHostFileVolumes obtains the hostFileVolume details from experiment CR spec
func getExperimentHostFileVolumes(clients ClientSets, expDetails *ExperimentDetails) ([]v1alpha1.HostFile, error) {
	chaosExperimentObj, err := clients.LitmusClient.LitmuschaosV1alpha1().ChaosExperiments(expDetails.Namespace).Get(expDetails.Name, metav1.GetOptions{})

    //var nilStructForHostFile v1alpha1.HostFile

	if err != nil {
		//return nilStructForHostFile, errors.Wrapf(err, "Unable to get ChaosExperiment Resource,  error: %v", err)
		return nil, errors.Wrapf(err, "Unable to get ChaosExperiment Resource,  error: %v", err)
	}
	expHostFileVolumes := chaosExperimentObj.Spec.Definition.HostFileVolumes

	return expHostFileVolumes, nil
}
