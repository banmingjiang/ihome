package models

func Upload(FileName string) {
	/*fdfsClient, err := NewFdfsClient("client.conf")
	if err != nil {
		t.Errorf("New FdfsClient error %s", err.Error())
		return
	}

	uploadResponse, err = fdfsClient.UploadByFilename("client.conf")
	if err != nil {
		t.Errorf("UploadByfilename error %s", err.Error())
	}
	t.Log(uploadResponse.GroupName)
	t.Log(uploadResponse.RemoteFileId)
	fdfsClient.DeleteFile(uploadResponse.RemoteFileId)
	*/
	// fdfsClient, dfs_err := fdfs_client.NewFdfsClient("conf/client.conf")
	// if dfs_err != nil {
	// 	beego.Info("use fdfs fail:", dfs_err.Error())
	// 	MakeLogs("use fdfs fail:", dfs_err.Error())
	// 	return
	// }
	// uploadResponse, upload_err := fdfsClient.UploadByFilename(FileName)
	// if upload_err != nil {
	// 	beego.Info("use UploadByFilename fail:", upload_err.Error())
	// 	MakeLogs("use UploadByFilename fail:", upload_err.Error())
	// }
	// beego.Info(uploadResponse.GroupName)
	// beego.Info(uploadResponse.RemoteFileId)

}
