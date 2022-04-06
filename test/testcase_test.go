package util
import (
    "testing"
    "fmt"
)

func Test_23803_backup_restore(t *testing.T) {

    var ret_val, ret_val1, ret_val2 bool

    //display etcd pod status before test
    ret_val, _ = oc_cli("export KUBECONFIG=/tmp/kubeconfig; oc get pod -n openshift-etcd -l app=etcd")
    if ret_val {
        fmt.Println("command execute succ.")
    } else {
        fmt.Println("command execute fail.")
    }
    //backup
    var nodeNameList []string
    nodeNameList = getNodeListByLabel("node-role.kubernetes.io/master=")
    master0 := nodeNameList[0]
    ret_val1, _ = oc_cli("export KUBECONFIG=/tmp/kubeconfig; oc debug node/" + master0 + " -- chroot /host /usr/local/bin/cluster-backup.sh /home/core/assets/backup")
    if ret_val1 {
        fmt.Println("backup command execute succ.")
    } else {
        fmt.Println("backup command execute fail.")
    }

    //restore    

    ret_val2, _ = oc_cli("export KUBECONFIG=/tmp/kubeconfig; oc debug node/" + master0 + " -- chroot /host /usr/local/bin/cluster-restore.sh /home/core/assets/backup")
    if ret_val2 {
        fmt.Println("restore command execute succ.")
    } else {
        fmt.Println("restore command execute fail.")
    }

    //check all etcd pod is Running status
    ret_val3 := etcdPodStatue()
    fmt.Printf("etcdPodStatue ret val %d:",ret_val3)
    if ret_val3 == 0 {
        fmt.Println("All etcd pods status are Running.")
    } else{
        fmt.Println("!!!etcd pods status are NOT Running.")
    }
    
    //delete etcd backup file on master node 
    delEtcdBackupDb(master0)    
}
