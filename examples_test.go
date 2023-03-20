package mountinfo_test

import (
	"fmt"
	"strings"

	"github.com/fntlnz/mountinfo"
)

func ExampleGetMountInfo() {
	minfo, _ := mountinfo.GetMountInfo("/proc/self/mountinfo")
	fmt.Printf("Mountpoint: %s", minfo[0].MountPoint)
}

func ExampleParseMountInfo() {
	lines := `26 25 0:24 / /sys/fs/cgroup/systemd rw,nosuid,nodev,noexec,relatime shared:9 - cgroup cgroup rw,xattr,release_agent=/usr/lib/systemd/systemd-cgroups-agent,name=systemd
515 24 0:3 net:[4026533140] /run/docker/netns/f46c0b2da189 rw shared:188 - nsfs nsfs rw
	`
	buf := strings.NewReader(lines)
	minfo, _ := mountinfo.ParseMountInfo(buf)
	fmt.Printf("Mountpoint 0: %s\n", minfo[0].MountPoint)
	fmt.Printf("Mountpoint 1: %s\n", minfo[1].MountPoint)
	fmt.Printf("MountSource 1: %s", minfo[1].MountSource)

	// Output: Mountpoint 0: /sys/fs/cgroup/systemd
	// Mountpoint 1: /run/docker/netns/f46c0b2da189
	// MountSource 1: nsfs
}
