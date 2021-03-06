// +build linux

package docker

import "os"

const (
	dockerSockPath = "/var/run/docker.sock"
)

func (r *Runner) exposeDocker() {
	r.exposeDockerEnv()
	if os.Getenv("DOCKER_HOST") == "" {
		sockPath := dockerSockPath
		if r.Task.Target.Ext != nil {
			serverSock, ok := r.Task.Target.Ext["server-socket"].(string)
			if ok && serverSock != "" {
				sockPath = serverSock
			}
		}
		r.Volumes = append(r.Volumes, sockPath+":"+sockPath)
	}
}

func currentUserIds() (uid, gid int, grps []int, err error) {
	uid = os.Getuid()
	gid = os.Getgid()
	grps, err = os.Getgroups()
	return
}
