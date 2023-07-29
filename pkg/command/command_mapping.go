package command

var CommandMapping = map[string]Command{
	"create_internet_cafe": new(CreateInternetCafe),
	"status":               new(StatusInternetCafe),
	"leave":                new(LeavePC),
	"book":                 new(AllocatePC),
}
