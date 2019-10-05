package client

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MySQLDeployment(domain string, password string) *appsv1.Deployment {
	/*
	***************************************************************
	************* Kube MySQL Deployment YAML **********************
	***************************************************************
	 */

	mysqlDep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: domain + "-mysql",
			Labels: map[string]string{
				"app": "wordpress",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "wordpress",
					"tier": "mysql",
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "wordpress",
						"tier": "mysql",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "mysql",
							Image: "mysql:5.6",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "mysql",
									ContainerPort: 3306,
								},
							},
							Env: []apiv1.EnvVar{
								apiv1.EnvVar{
									Name:  "MYSQL_ROOT_PASSWORD",
									Value: password,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								apiv1.VolumeMount{
									Name:      domain + "-mysql-persistent-storage",
									MountPath: "/var/lib/mysql/" + domain,
								},
							},
						},
					},
					Volumes: []apiv1.Volume{
						apiv1.Volume{
							Name: domain + "-mysql-persistent-storage",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: domain + "-mysql-pv",
								},
							},
						},
					},
				},
			},
		},
	}
	return mysqlDep
}

///////////////////////////////////////////////////////

func MySQLService(domain string) *apiv1.Service {
	/*
	***************************************************************
	************* Kube MySQL Service YAML *********************
	***************************************************************
	 */
	mysqlServ := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: domain + "-mysql",
			Labels: map[string]string{
				"app": domain + "-wordpress",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port: 3306,
				},
			},
			Selector: map[string]string{
				"app":  domain + "-wordpress",
				"tier": "mysql",
			},
		},
	}
	return mysqlServ
}

///////////////////////////////////////////////////////

func MySQLPV(domain string) *apiv1.PersistentVolume {
	/*
	***************************************************************
	************* Kube WordPress Persistent Volume YAML ***********
	***************************************************************
	 */
	wordPV := &apiv1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: domain + "-pv",
			Labels: map[string]string{
				"type": "local",
			},
		},
		Spec: apiv1.PersistentVolumeSpec{
			StorageClassName: "manual",
			Capacity: apiv1.ResourceList{
				apiv1.ResourceName(apiv1.ResourceStorage): resource.MustParse("1Gi"),
			},
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteOnce,
			},
			PersistentVolumeSource: apiv1.PersistentVolumeSource{
				HostPath: &apiv1.HostPathVolumeSource{
					Path: "/mnt/" + domain,
				},
			},
		},
	}
	return wordPV

}

///////////////////////////////////////////////////////

func MySQLPVC(domain string) *apiv1.PersistentVolumeClaim {
	/*
	***************************************************************
	************* Kube MySQL Persistent Volume Claim YAML *********
	***************************************************************
	 */
	mysqlPVClaim := &apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: domain + "-mysql-pv",
			Labels: map[string]string{
				"app": "wordpress",
			},
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.ReadWriteOnce,
			},
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					apiv1.ResourceName(apiv1.ResourceStorage): resource.MustParse("1Gi"),
				},
			},
		},
	}

	return mysqlPVClaim
}
