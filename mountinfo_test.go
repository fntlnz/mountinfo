package mountinfo

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type ParseMountData struct {
	rawline     string
	expectedset Mountinfo
}

// TestParseMountString data set, please add more cases if you feel
func ParseMountDataset() []ParseMountData {
	return []ParseMountData{
		{
			rawline: "515 24 0:3 net:[4026533140] /run/docker/netns/f46c0b2da189 rw shared:188 - nsfs nsfs rw",
			expectedset: Mountinfo{
				MountID:        "515",
				ParentID:       "24",
				MajorMinor:     "0:3",
				Root:           "net:[4026533140]",
				MountPoint:     "/run/docker/netns/f46c0b2da189",
				MountOptions:   "rw",
				OptionalFields: "shared:188",
				FilesystemType: "nsfs",
				MountSource:    "nsfs",
				SuperOptions:   "rw",
			},
		},
		{
			rawline: "26 25 0:24 / /sys/fs/cgroup/systemd rw,nosuid,nodev,noexec,relatime shared:9 - cgroup cgroup rw,xattr,release_agent=/usr/lib/systemd/systemd-cgroups-agent,name=systemd",
			expectedset: Mountinfo{
				MountID:        "26",
				ParentID:       "25",
				MajorMinor:     "0:24",
				Root:           "/",
				MountPoint:     "/sys/fs/cgroup/systemd",
				MountOptions:   "rw,nosuid,nodev,noexec,relatime",
				OptionalFields: "shared:9",
				FilesystemType: "cgroup",
				MountSource:    "cgroup",
				SuperOptions:   "rw,xattr,release_agent=/usr/lib/systemd/systemd-cgroups-agent,name=systemd",
			},
		},
		{
			rawline: "824 723 0:52 / /var/lib/containers/storage/overlay/62ef728cd5abf2cf711bed7912828d1283ca2b5cb2200e65941d86e22cd6c205/merged rw,nodev,relatime - overlay overlay rw,context=\"system_u:object_r:container_file_t:s0:c120,c334\",lowerdir=/var/lib/containers/storage/overlay/l/NZWCWDRTRJEXYQAOIXQLIOWOOJ:/var/lib/containers/storage/overlay/l/EG2LSO3NQSV6ZKSAG2SF7BUO4L:/var/lib/containers/storage/overlay/l/EMD4ZHWPE6MJ4KZBHPTFDVJV7B:/var/lib/containers/storage/overlay/l/MYIZJQPW4OD3J4TEJRXTMV5NVJ:/var/lib/containers/storage/overlay/l/6MMB3RGZNRYW6YZQI3FUUN2GTM:/var/lib/containers/storage/overlay/l/O5Z3Y63L3AVQLPB7WND73PQG6W:/var/lib/containers/storage/overlay/l/UACPLH6JBVALN7TSY7AYPRQT4F:/var/lib/containers/storage/overlay/l/CT6LZD32636MU36BUPS4KQ7GUK:/var/lib/containers/storage/overlay/l/S7ZQGTZDQOU5AH2E3R74EPBKW7:/var/lib/containers/storage/overlay/l/66SOAQGJFJDCHQ6G3WVQENMBEP:/var/lib/containers/storage/overlay/l/RQQO2I2OI47TML4V3DPSMMQ5KQ:/var/lib/containers/storage/overlay/l/BH2MWLJXSNYW4NRNQZHVJVF454:/var/lib/containers/storage/overlay/l/L7WEQU6QPK2GQJZUFPBHVEOBFM,upperdir=/var/lib/containers/storage/overlay/62ef728cd5abf2cf711bed7912828d1283ca2b5cb2200e65941d86e22cd6c205/diff,workdir=/var/lib/containers/storage/overlay/62ef728cd5abf2cf711bed7912828d1283ca2b5cb2200e65941d86e22cd6c205/work,metacopy=on",
			expectedset: Mountinfo{
				MountID:        "824",
				ParentID:       "723",
				MajorMinor:     "0:52",
				Root:           "/",
				MountPoint:     "/var/lib/containers/storage/overlay/62ef728cd5abf2cf711bed7912828d1283ca2b5cb2200e65941d86e22cd6c205/merged",
				MountOptions:   "rw,nodev,relatime",
				OptionalFields: "",
				FilesystemType: "overlay",
				MountSource:    "overlay",
				SuperOptions:   "rw,context=\"system_u:object_r:container_file_t:s0:c120,c334\",lowerdir=/var/lib/containers/storage/overlay/l/NZWCWDRTRJEXYQAOIXQLIOWOOJ:/var/lib/containers/storage/overlay/l/EG2LSO3NQSV6ZKSAG2SF7BUO4L:/var/lib/containers/storage/overlay/l/EMD4ZHWPE6MJ4KZBHPTFDVJV7B:/var/lib/containers/storage/overlay/l/MYIZJQPW4OD3J4TEJRXTMV5NVJ:/var/lib/containers/storage/overlay/l/6MMB3RGZNRYW6YZQI3FUUN2GTM:/var/lib/containers/storage/overlay/l/O5Z3Y63L3AVQLPB7WND73PQG6W:/var/lib/containers/storage/overlay/l/UACPLH6JBVALN7TSY7AYPRQT4F:/var/lib/containers/storage/overlay/l/CT6LZD32636MU36BUPS4KQ7GUK:/var/lib/containers/storage/overlay/l/S7ZQGTZDQOU5AH2E3R74EPBKW7:/var/lib/containers/storage/overlay/l/66SOAQGJFJDCHQ6G3WVQENMBEP:/var/lib/containers/storage/overlay/l/RQQO2I2OI47TML4V3DPSMMQ5KQ:/var/lib/containers/storage/overlay/l/BH2MWLJXSNYW4NRNQZHVJVF454:/var/lib/containers/storage/overlay/l/L7WEQU6QPK2GQJZUFPBHVEOBFM,upperdir=/var/lib/containers/storage/overlay/62ef728cd5abf2cf711bed7912828d1283ca2b5cb2200e65941d86e22cd6c205/diff,workdir=/var/lib/containers/storage/overlay/62ef728cd5abf2cf711bed7912828d1283ca2b5cb2200e65941d86e22cd6c205/work,metacopy=on",
			},
		},
	}
}

func TestParseMountString(t *testing.T) {
	for _, e := range ParseMountDataset() {
		info := ParseMountInfoString(e.rawline)

		if reflect.DeepEqual(e.expectedset, *info) == false {
			t.Errorf("expected %v got %v", e.expectedset, *info)
		}
	}
}

func TestParseMountInfo(t *testing.T) {
	buf := bytes.Buffer{}
	expectedSet := []Mountinfo{}
	for _, e := range ParseMountDataset() {
		buf.WriteString(fmt.Sprintf("%s\n", e.rawline))
		expectedSet = append(expectedSet, e.expectedset)
	}

	info, err := ParseMountInfo(&buf)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if reflect.DeepEqual(expectedSet, info) == false {
		t.Errorf("expected %v got %v", expectedSet, info)
	}
}
