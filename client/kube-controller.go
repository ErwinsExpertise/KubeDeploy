package client

import (
	"log"
	"os"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ClientInit() *kubernetes.Clientset {
	/*
	***************************************************************
	************* Find .kube config and assign to variable ********
	***************************************************************
	 */
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Println(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
	}
	return clientset
}

func CreateSite(domain string, password string) string {
	clientset := ClientInit()
	/*
	***************************************************************
	************* Create Kube Client to connect to API  ***********
	***************************************************************
	 */
	depClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	servClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	pvClient := clientset.CoreV1().PersistentVolumeClaims(apiv1.NamespaceDefault)
	//pevClient := clientset.CoreV1().PersistentVolumes()

	/*
	***************************************************************
	************* Attempt to Create and return message ************
	***************************************************************
	 */
	/*
		pevResult, err := pevClient.Create(wordPV)
		if err != nil {
			log.Print(err)
		}
	*/

	/*
		//Let's create MySQL First
		mysqldepResult, err := depClient.Create(mysqlDep)
		if err != nil {
			log.Print(err)
		}
		mysqlservResult, err := servClient.Create(mysqlServ)
		if err != nil {
			log.Print(err)
		}
		mysqlpvResult, err := pvClient.Create(mysqlPVClaim)
		if err != nil {
			log.Print(err)
		}
	*/

	//Now we deploy WordPress
	depResult, err := depClient.Create(WordPressDeployment(domain))
	if err != nil {
		log.Print(err)
	}
	servResult, err := servClient.Create(WordPressService(domain))
	if err != nil {
		log.Print(err)
	}
	pvResult, err := pvClient.Create(WordPressPVC(domain))
	if err != nil {
		log.Print(err)
	}

	result := (" WordPress Deployment: " + depResult.GetObjectMeta().GetName() +
		" WordPress Service: " + servResult.GetObjectMeta().GetName() +
		" WordPress PV Claim: " + pvResult.GetObjectMeta().GetName())
	return result
}

func GetDeployments() []string {
	client := ClientInit()
	var sites []string
	deps, err := client.ExtensionsV1beta1().Deployments(meta_v1.NamespaceDefault).List(meta_v1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for _, dep := range deps.Items {
		sites = append(sites, dep.Name)
	}

	return sites
}

func int32Ptr(i int32) *int32 { return &i }
