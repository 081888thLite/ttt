package ttt

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFeed(t *testing.T) {
	msg := "Received\n"
	msgAfterReading := "Received"
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, msg)
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}

	fed, _ := feed(in)
	if fed != msgAfterReading {
		t.Errorf("Expected %v,\n got: %v", msgAfterReading, fed)
	}
}

func TestSeed(t *testing.T) {
	msg := "Sent"
	out, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	_, err = io.WriteString(out, msg)
	if err != nil {
		t.Fatal(err)
	}

	_, err = out.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}
	_, err = seed(out, msg)
	if err != nil {
		t.Errorf("Expected no errors when sending '%v',\n got: %v", msg, err)
	}
}

func TestGetLastRead(t *testing.T) {
	msgStatus := MsgStatus{Msg: "Latest"}
	client := StubClient{LastRead: msgStatus}
	accessed := client.GetLastRead()
	if accessed != "Latest" {
		t.Error("Expected the client's LastRead.Msg as string,\n got:", accessed)
	}
}
func TestGetLastSent(t *testing.T) {
	msgStatus := MsgStatus{Msg: "Latest"}
	client := StubClient{LastWrote: msgStatus}
	accessed := client.GetLastSent()
	if accessed != "Latest" {
		t.Error("Expected the client's LastWrote.Msg as string,\n got:", accessed)
	}
}

func Test_feed(t *testing.T) {
	type args struct {
		source *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test feed stdIn Default",
			args:    args{nil},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := feed(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("feed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("feed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_seed(t *testing.T) {
	type args struct {
		source *os.File
		msg    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test seed stdOut Default",
			args:    args{nil, "Seeded"},
			want:    "Seeded",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := seed(tt.args.source, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("seed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("seed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSys_Write(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Writes to stdOut",
			fields: fields{MsgStatus{"Wrote", nil}, MsgStatus{"", nil}},
			args:   args{"Wrote"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Sys{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			client.Write(tt.args.msg)
		})
	}
}

func TestSys_Read(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Sys Read works",
			fields: fields{MsgStatus{"", nil}, MsgStatus{"", nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Sys{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			client.Read()
		})
	}
}

func TestStubClient_Write(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Sets LastWrote",
			fields: fields{MsgStatus{"Wrote", nil}, MsgStatus{"", nil}},
			args:   args{"Wrote"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &StubClient{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			client.Write(tt.args.msg)
		})
	}
}

func TestSys_GetLastRead(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Returns last read from stdIn",
			fields: fields{MsgStatus{"", nil}, MsgStatus{"was Read", nil}},
			want:   "was Read",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Sys{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			if got := client.GetLastRead(); got != tt.want {
				t.Errorf("Sys.GetLastRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStubClient_Read(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Reads whatever is in the LastRead field",
			fields: fields{MsgStatus{"", nil}, MsgStatus{"LastRead", nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &StubClient{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			client.Read()
		})
	}
}

func TestStubClient_GetLastRead(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Reads whatever is in the LastRead field",
			fields: fields{MsgStatus{"", nil}, MsgStatus{"LastRead", nil}},
			want:   "LastRead",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &StubClient{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			if got := client.GetLastRead(); got != tt.want {
				t.Errorf("StubClient.GetLastRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSys_GetLastSent(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Returns LastWrote to stdOut",
			fields: fields{MsgStatus{"LastWrote", nil}, MsgStatus{"", nil}},
			want:   "LastWrote",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Sys{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			if got := client.GetLastSent(); got != tt.want {
				t.Errorf("Sys.GetLastSent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStubClient_GetLastSent(t *testing.T) {
	type fields struct {
		LastWrote MsgStatus
		LastRead  MsgStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Returns whatever is in the LastSent field",
			fields: fields{MsgStatus{"LastSent", nil}, MsgStatus{"LastRead", nil}},
			want:   "LastSent",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &StubClient{
				LastWrote: tt.fields.LastWrote,
				LastRead:  tt.fields.LastRead,
			}
			if got := client.GetLastSent(); got != tt.want {
				t.Errorf("StubClient.GetLastSent() = %v, want %v", got, tt.want)
			}
		})
	}
}
