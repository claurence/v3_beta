package models

import "time"

type V3Error struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
	Code   int    `json:"code"`
}

type V2Error struct {
	Description string `json:"description,omitempty"`
	ErrorCode   string `json:"error_code,omitempty"`
	Code        int    `json:"code,omitempty"`
}

type V3AppModel struct {
	Errors     []V3Error `json:"errors"`
	Name       string
	Guid       string
	Error_Code string
	Processes  string
	Instances  int `json:"total_desired_instances"`
}

type V3IsolationSegmentsModel struct {
	IsoSegs []V3IsolationSegmentModel `json:"resources"`
}

type V3IsolationSegmentModel struct {
	Guid      string
	Name      string
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type V3ProcessModel struct {
	Type      string
	Instances int
	Memory    int        `json:"memory_in_mb"`
	Disk      int        `json:"disk_in_mb"`
	Links     LinksModel `json:"links"`
}

type V3TaskModel struct {
	Id        int       `json:"sequence_id"`
	Name      string    `json:"name"`
	Guid      string    `json:"guid"`
	Command   string    `json:"command"`
	State     string    `json:"state"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type V3AppsModel struct {
	Errors []V3Error    `json:"errors"`
	Apps   []V3AppModel `json:"resources"`
}

type V3ProcessesModel struct {
	Processes []V3ProcessModel `json:"resources"`
}

type V3TasksModel struct {
	Tasks []V3TaskModel `json:"resources"`
}

type LinkModel struct {
	Href string
}

type LinksModel struct {
	App   LinkModel
	Space LinkModel
}

type RelationshipModel struct {
	Data []map[string]string
}

type V3PackageModel struct {
	Errors     []V3Error `json:"errors"`
	Guid       string
	Error_Code string
}

type V3DropletModel struct {
	Errors []V3Error `json:"errors"`
	Guid   string
}

type V3BuildModel struct {
	Errors  []V3Error `json:"errors"`
	Guid    string
	Droplet V3DropletModel `json:"droplet"`
}

type MetadataModel struct {
	Guid string `json:"guid"`
}

type EntityModel struct {
	Name string `json:"name"`
}
type RouteEntityModel struct {
	Host string `json:"host"`
}

type DomainsModel struct {
	V2Error
	NextUrl   string        `json:"next_url,omitempty"`
	Resources []DomainModel `json:"resources"`
}
type DomainModel struct {
	Metadata MetadataModel `json:"metadata"`
	Entity   EntityModel   `json:"entity"`
}

type OrgsModel struct {
	Orgs []OrgModel `json:"resources"`
}

type OrgModel struct {
	Metadata MetadataModel `json:"metadata"`
	Entity   EntityModel   `json:"entity"`
}

type RouteModel struct {
	V2Error
	Metadata MetadataModel    `json:"metadata"`
	Entity   RouteEntityModel `json:"entity"`
}
type RoutesModel struct {
	V2Error
	Routes []RouteModel `json:"resources"`
}
