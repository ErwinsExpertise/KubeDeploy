package client

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func WordPressDeployment(domain string) *appsv1.Deployment {
	var conf Database
	conf.BuildConfig()
	/*
	***************************************************************
	************* Kube WordPress Deployment YAML ******************
	***************************************************************
	 */
	wordDep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: domain + "-wordpress",
			Labels: map[string]string{
				"app": domain,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  domain,
					"tier": "frontend",
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  domain,
						"tier": "frontend",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  domain,
							Image: "wordpress:latest",
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
							Env: []apiv1.EnvVar{
								apiv1.EnvVar{
									Name:  "WORDPRESS_DB_HOST",
									Value: conf.Host,
								},
								apiv1.EnvVar{
									Name:  "WORDPRESS_DB_USER",
									Value: conf.Username,
								},
								apiv1.EnvVar{
									Name:  "WORDPRESS_DB_PASSWORD",
									Value: conf.Password,
								},
								apiv1.EnvVar{
									Name:  "WORDPRESS_DB_NAME",
									Value: domain,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								apiv1.VolumeMount{
									Name:      domain + "-persistent-storage",
									MountPath: "/var/www/" + domain,
								},
							},
						},
					},
					Volumes: []apiv1.Volume{
						apiv1.Volume{
							Name: domain + "-persistent-storage",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: domain + "-pv",
								},
							},
						},
					},
				},
			},
		},
	}
	return wordDep

}

func WordPressService(domain string) *apiv1.Service {
	/*
	***************************************************************
	************* Kube WordPress Service YAML *********************
	***************************************************************
	 */
	wordServ := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: domain + "-wordpress",
			Labels: map[string]string{
				"app": domain,
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port: 80,
				},
			},
			Selector: map[string]string{
				"app":  domain,
				"tier": "frontend",
			},
			Type: "LoadBalancer",
		},
	}
	return wordServ
}

func WordPressPVC(domain string) *apiv1.PersistentVolumeClaim {
	/*
	***************************************************************
	************* Kube WordPress Persistent Volume Claim YAML *****
	***************************************************************
	 */
	var conf Database
	conf.BuildConfig()
	wordPVClaim := &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: domain + "-pv",
			Labels: map[string]string{
				"app": domain,
			},
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteMany,
			},
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					apiv1.ResourceName(apiv1.ResourceStorage): resource.MustParse("1Gi"),
				},
			},
			StorageClassName: conf.StorageName,
		},
	}
	return wordPVClaim
}

