package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"strconv"
	"strings"
	"time"
)

type Vpc struct {
	Name      string
	Id        string
	CidrBlock string
}

type Vpcs []Vpc

func (v Vpcs) Len() int {
	return len(v)
}

func (v Vpcs) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v Vpcs) Less(i, j int) bool {
	currentParts := strings.Split(v[i].CidrBlock, "/")
	currentIp := currentParts[0]
	currentPrefix := currentParts[1]
	currentIpParts := strings.Split(currentIp, ".")

	nextParts := strings.Split(v[j].CidrBlock, "/")
	nextIp := nextParts[0]
	nextPrefix := currentParts[1]
	nextIpParts := strings.Split(nextIp, ".")

	oct1 := v.intCompare(currentIpParts[0], nextIpParts[0])
	oct2 := v.intCompare(currentIpParts[1], nextIpParts[1])
	oct3 := v.intCompare(currentIpParts[2], nextIpParts[2])
	oct4 := v.intCompare(currentIpParts[3], nextIpParts[3])
	prefix := v.intCompare(currentPrefix, nextPrefix)

	if oct1 != 0 {
		return oct1 < 0
	}
	if oct2 != 0 {
		return oct2 < 0
	}
	if oct3 != 0 {
		return oct3 < 0
	}
	if oct4 != 0 {
		return oct4 < 0
	}
	if prefix != 0 {
		return prefix < 0
	}

	return false
}

func (v Vpcs) intCompare(i, j string) int {
	a, _ := strconv.Atoi(i)
	b, _ := strconv.Atoi(j)

	if a < b {
		return -1 // a is less than b
	} else if a > b {
		return 1 // a is greater than b
	}
	return 0 // a is equal to b
}

type MyInstanceInfo struct {
	Name       string
	Id         string
	Type       string
	State      string
	Ami        string
	LaunchTime time.Time
	PrivateIp  string
	PublicIp   string
}

type Image struct {
	Name         string
	ImageId      string
	ImageType    types.ImageTypeValues
	CreationDate time.Time
	State        types.ImageState
}

func NewImage(other types.Image) Image {
	creationTime, _ := time.Parse("2006-01-02T15:04:05.000Z", *other.CreationDate)
	return Image{
		Name:         *other.Name,
		ImageId:      *other.ImageId,
		ImageType:    other.ImageType,
		CreationDate: creationTime,
		State:        other.State,
	}
}
