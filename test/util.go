package util

import (
    "fmt"
    "strings"
    "bytes"
    "os/exec"
    "log"
)

func oc_cli(commandName string) (bool, string){

    cmd := exec.Command("/bin/sh", "-c", commandName)
    fmt.Printf("execute oc command line: \n%s\n", cmd.Args)

    var ret_val = true
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout 
    cmd.Stderr = &stderr 
    err := cmd.Run()
    outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
    fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
        ret_val = false
        return ret_val, errStr
    }
    return ret_val, outStr
}
     
func getNodeListByLabel(labelKey string) []string {
        var ret_val=true
        var outStr=""   
        ret_val, outStr = oc_cli("export KUBECONFIG=/tmp/kubeconfig;" + "oc get node -l " + labelKey + " -o=jsonpath={.items[*].metadata.name}")
        if ret_val == false {
           //    exit 
        }
        nodeNameList := strings.Fields(outStr)
        return nodeNameList
}

func getPodListByLabel(labelKey string) []string {
        var ret_val=true
        var outStr=""
	ret_val, outStr = oc_cli("export KUBECONFIG=/tmp/kubeconfig; " + "oc get pod -n openshift-etcd -l " + labelKey + " -o=jsonpath={.items[*].metadata.name}")
        if ret_val == false {
                  //     exit 
        }
	podNameList := strings.Fields(outStr)
	return podNameList
}

func etcdPodStatue() int {
    var commandName = "export KUBECONFIG=/tmp/kubeconfig; oc get pod -n openshift-etcd -l \"app=etcd\" |awk '{print $3}'|awk 'NR>2{print line}{line=$0} END{print line}'|grep -v Running"

    outstr, errstr := runCmdOnMaster(commandName)
    fmt.Printf("**etcdPodStatue**: out:\n%s\nerr:\n%s\n length of outstr:%d\n", outstr, errstr, len(outstr)) 
    
    if outstr == "" { //grep non-Running is null, mean etcd pods are Running status
        return 0
    }else{
        fmt.Printf("errstr in etcdPodStatue: %s\n",errstr)
        return 1
    }    
}

func runCmdOnMaster(cmdName string) (string, string) {
    cmd := exec.Command("/bin/sh", "-c", cmdName)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    outstr, errstr := string(stdout.Bytes()), string(stderr.Bytes())
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    return outstr, errstr
}

func delEtcdBackupDb(master string)  {
    var commandName = "export KUBECONFIG=/tmp/kubeconfig; oc debug node/" + master + " -- chroot /host rm -rf /home/core/assets/backup"   
    outstr, errstr := runCmdOnMaster(commandName) 
    fmt.Printf("out:\n%s\nerr:\n%s\n", outstr, errstr)    
}
