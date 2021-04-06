package kgo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestFile_GetExt(t *testing.T) {
	var ext string

	ext = KFile.GetExt(fileGo)
	assert.Equal(t, "go", ext)

	ext = KFile.GetExt(fileGitkee)
	assert.Equal(t, "gitkeep", ext)

	ext = KFile.GetExt(fileSongs)
	assert.Equal(t, "txt", ext)

	ext = KFile.GetExt(fileNone)
	assert.Empty(t, ext)
}

func BenchmarkFile_GetExt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetExt(fileMd)
	}
}

func TestFile_ReadFile(t *testing.T) {
	var bs []byte
	var err error

	bs, err = KFile.ReadFile(fileMd)
	assert.NotEmpty(t, bs)
	assert.Nil(t, err)

	//不存在的文件
	bs, err = KFile.ReadFile(fileNone)
	assert.NotNil(t, err)
}

func BenchmarkFile_ReadFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ReadFile(fileMd)
	}
}

func TestFile_ReadInArray(t *testing.T) {
	var sl []string
	var err error

	sl, err = KFile.ReadInArray(fileDante)
	assert.Equal(t, 19568, len(sl))

	//不存在的文件
	sl, err = KFile.ReadInArray(fileNone)
	assert.NotNil(t, err)
}

func BenchmarkFile_ReadInArray(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ReadInArray(fileMd)
	}
}

func TestFile_ReadFirstLine(t *testing.T) {
	var res []byte

	res = KFile.ReadFirstLine(fileDante)
	assert.NotEmpty(t, res)

	//不存在的文件
	res = KFile.ReadFirstLine(fileNone)
	assert.Empty(t, res)
}

func BenchmarkFile_ReadFirstLine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.ReadFirstLine(fileMd)
	}
}

func TestFile_ReadLastLine(t *testing.T) {
	var res []byte

	res = KFile.ReadLastLine(changLog)
	assert.NotEmpty(t, res)

	//不存在的文件
	res = KFile.ReadLastLine(fileNone)
	assert.Empty(t, res)
}

func BenchmarkFile_ReadLastLine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.ReadLastLine(fileMd)
	}
}

func TestFile_WriteFile(t *testing.T) {
	var err error

	err = KFile.WriteFile(putfile, bytsHello)
	assert.Nil(t, err)

	//设置权限
	err = KFile.WriteFile(putfile, bytsHello, 0777)
	assert.Nil(t, err)

	//无权限写
	err = KFile.WriteFile(rootFile1, bytsHello, 0777)
	if KOS.IsLinux() || KOS.IsMac() {
		assert.NotNil(t, err)
	}
}

func BenchmarkFile_WriteFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filename := fmt.Sprintf("./testdata/file/putfile_%d", i)
		_ = KFile.WriteFile(filename, bytsHello)
	}
}

func TestFile_AppendFile(t *testing.T) {
	var err error

	//创建
	err = KFile.AppendFile(apndfile, bytsHello)
	assert.Nil(t, err)

	//追加
	err = KFile.AppendFile(apndfile, bytsHello)
	assert.Nil(t, err)

	//空路径
	err = KFile.AppendFile("", bytsHello)
	assert.NotNil(t, err)

	//权限不足
	err = KFile.AppendFile(rootFile1, bytsHello)
	if KOS.IsLinux() || KOS.IsMac() {
		assert.NotNil(t, err)
	}
}

func BenchmarkFile_AppendFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KFile.AppendFile(apndfile, bytsHello)
	}
}

func TestFile_GetMime(t *testing.T) {
	var res string

	res = KFile.GetMime(imgPng, false)
	assert.NotEmpty(t, res)

	res = KFile.GetMime(fileDante, true)
	if KOS.IsWindows() {
		assert.NotEmpty(t, res)
	}

	//不存在的文件
	res = KFile.GetMime(fileNone, true)
	assert.Empty(t, res)
}

func BenchmarkFile_GetMime_Fast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetMime(fileMd, true)
	}
}

func BenchmarkFile_GetMime_NoFast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetMime(fileMd, false)
	}
}

func TestFile_FileSize(t *testing.T) {
	var res int64

	res = KFile.FileSize(changLog)
	assert.Greater(t, res, int64(0))

	//不存在的文件
	res = KFile.FileSize(fileNone)
	assert.Equal(t, int64(-1), res)
}

func BenchmarkFile_FileSize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FileSize(fileMd)
	}
}

func TestFile_DirSize(t *testing.T) {
	var res int64

	res = KFile.DirSize(dirCurr)
	assert.Greater(t, res, int64(0))

	//不存在的目录
	res = KFile.DirSize(fileNone)
	assert.Equal(t, int64(0), res)
}

func BenchmarkFile_DirSize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.DirSize(dirTdat)
	}
}

func TestFile_IsExist(t *testing.T) {
	var res bool

	res = KFile.IsExist(changLog)
	assert.True(t, res)

	res = KFile.IsExist(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsExist(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsExist(fileMd)
	}
}

func TestFile_IsReadable(t *testing.T) {
	var res bool

	res = KFile.IsReadable(dirTdat)
	assert.True(t, res)

	//不存在的目录
	res = KFile.IsReadable(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsReadable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsReadable(dirTdat)
	}
}

func TestFile_IsWritable(t *testing.T) {
	var res bool

	res = KFile.IsWritable(dirTdat)
	assert.True(t, res)

	//不存在的目录
	res = KFile.IsWritable(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsWritable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsWritable(dirTdat)
	}
}

func TestFile_IsExecutable(t *testing.T) {
	var res bool

	res = KFile.IsExecutable(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsExecutable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsExecutable(fileMd)
	}
}

func TestFile_IsLink(t *testing.T) {
	//创建链接文件
	if !KFile.IsExist(fileLink) {
		_ = os.Symlink(filePubPem, fileLink)
	}

	var res bool

	res = KFile.IsLink(fileLink)
	assert.True(t, res)

	res = KFile.IsLink(changLog)
	assert.False(t, res)
}

func BenchmarkFile_IsLink(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsLink(fileLink)
	}
}

func TestFile_IsFile(t *testing.T) {
	tests := []struct {
		f        string
		t        LkkFileType
		expected bool
	}{
		{"", FILE_TYPE_ANY, false},
		{fileNone, FILE_TYPE_ANY, false},
		{fileGo, FILE_TYPE_ANY, true},
		{fileMd, FILE_TYPE_LINK, false},
		{fileLink, FILE_TYPE_LINK, true},
		{fileLink, FILE_TYPE_REGULAR, false},
		{fileGitkee, FILE_TYPE_REGULAR, true},
		{fileLink, FILE_TYPE_COMMON, true},
		{imgJpg, FILE_TYPE_COMMON, true},
	}
	for _, test := range tests {
		actual := KFile.IsFile(test.f, test.t)
		assert.Equal(t, test.expected, actual)
	}
}

func BenchmarkFile_IsFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsFile(fileMd, FILE_TYPE_ANY)
	}
}

func TestFile_IsDir(t *testing.T) {
	var res bool

	res = KFile.IsDir(fileMd)
	assert.False(t, res)

	res = KFile.IsDir(fileNone)
	assert.False(t, res)

	res = KFile.IsDir(dirTdat)
	assert.True(t, res)
}

func BenchmarkFile_IsDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsDir(dirTdat)
	}
}

func TestFile_IsBinary(t *testing.T) {
	var res bool
	res = KFile.IsBinary(changLog)
	assert.False(t, res)

	//TODO true
}

func BenchmarkFile_IsBinary(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsBinary(changLog)
	}
}

func TestFile_IsImg(t *testing.T) {
	var res bool

	res = KFile.IsImg(fileMd)
	assert.False(t, res)

	res = KFile.IsImg(imgSvg)
	assert.True(t, res)

	res = KFile.IsImg(imgPng)
	assert.True(t, res)
}

func BenchmarkFile_IsImg(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsImg(imgPng)
	}
}

func TestFile_Mkdir(t *testing.T) {
	var err error

	err = KFile.Mkdir(dirNew, 0777)
	assert.Nil(t, err)
}

func BenchmarkFile_Mkdir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dname := fmt.Sprintf(dirNew+"/tmp_%d", i)
		_ = KFile.Mkdir(dname, 0777)
	}
}

func TestFile_AbsPath(t *testing.T) {
	var res string

	res = KFile.AbsPath(changLog)
	assert.NotEqual(t, '.', rune(res[0]))

	res = KFile.AbsPath(fileNone)
	assert.NotEmpty(t, res)
}

func BenchmarkFile_AbsPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.AbsPath(changLog)
	}
}

func TestFile_RealPath(t *testing.T) {
	var res string

	res = KFile.RealPath(fileMd)
	assert.NotEmpty(t, res)

	res = KFile.RealPath(fileNone)
	assert.Empty(t, res)
}

func BenchmarkFile_RealPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.RealPath(fileMd)
	}
}

func TestFile_TouchRenameUnlink(t *testing.T) {
	var res bool
	var err error

	res = KFile.Touch(touchfile, 2097152)
	assert.True(t, res)

	err = KFile.Rename(touchfile, renamefile)
	assert.Nil(t, err)

	err = KFile.Unlink(renamefile)
	assert.Nil(t, err)
}

func BenchmarkFile_Touch(b *testing.B) {
	b.ResetTimer()
	var filename string
	for i := 0; i < b.N; i++ {
		filename = fmt.Sprintf(dirTouch+"/zero_%d", i)
		KFile.Touch(filename, 0)
	}
}

func BenchmarkFile_Rename(b *testing.B) {
	b.ResetTimer()
	var f1, f2 string
	for i := 0; i < b.N; i++ {
		f1 = fmt.Sprintf(dirTouch+"/zero_%d", i)
		f2 = fmt.Sprintf(dirTouch+"/zero_re%d", i)
		_ = KFile.Rename(f1, f2)
	}
}

func BenchmarkFile_Unlink(b *testing.B) {
	b.ResetTimer()
	var filename string
	for i := 0; i < b.N; i++ {
		filename = fmt.Sprintf(dirTouch+"/zero_re%d", i)
		_ = KFile.Unlink(filename)
	}
}

func TestFile_CopyFile(t *testing.T) {
	var res int64
	var err error

	//忽略已存在的
	res, err = KFile.CopyFile(imgPng, imgCopy, FILE_COVER_IGNORE)
	assert.Nil(t, err)

	//覆盖已存在的
	res, err = KFile.CopyFile(imgPng, imgCopy, FILE_COVER_ALLOW)
	assert.Greater(t, res, int64(0))

	//禁止覆盖
	res, err = KFile.CopyFile(imgPng, imgCopy, FILE_COVER_DENY)
	assert.NotNil(t, err)

	//源和目标文件相同
	res, err = KFile.CopyFile(imgPng, imgPng, FILE_COVER_ALLOW)
	assert.Equal(t, int64(0), res)
	assert.Nil(t, err)

	//拷贝大文件
	KFile.Touch(touchfile, 2097152)
	res, err = KFile.CopyFile(touchfile, copyfile, FILE_COVER_ALLOW)

	//目标为空
	res, err = KFile.CopyFile(imgPng, "", FILE_COVER_ALLOW)
	assert.NotNil(t, err)

	//源非正常文件
	res, err = KFile.CopyFile(".", "", FILE_COVER_ALLOW)
	assert.NotNil(t, err)
}

func BenchmarkFile_CopyFile(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/diglett_copy_%d.png", i)
		_, _ = KFile.CopyFile(imgPng, des, FILE_COVER_ALLOW)
	}
}

func TestFile_FastCopy(t *testing.T) {
	var res int64
	var err error

	res, err = KFile.FastCopy(imgJpg, fastcopyfile)
	assert.Greater(t, res, int64(0))

	//源文件不存在
	res, err = KFile.FastCopy(fileNone, fastcopyfile)
	assert.NotNil(t, err)

	//目标为空
	res, err = KFile.FastCopy(imgJpg, "")
	assert.NotNil(t, err)
}

func BenchmarkFile_FastCopy(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/fast_copy_%d", i)
		_, _ = KFile.FastCopy(imgJpg, des)
	}
}

func TestFile_CopyLink(t *testing.T) {
	var err error

	//源和目标相同
	err = KFile.CopyLink(fileLink, fileLink)
	assert.Nil(t, err)

	err = KFile.CopyLink(fileLink, copyLink)
	assert.Nil(t, err)

	//源文件不存在
	err = KFile.CopyLink(fileNone, copyLink)
	assert.NotNil(t, err)

	//目标为空
	err = KFile.CopyLink(fileLink, "")
	assert.NotNil(t, err)
}

func BenchmarkFile_CopyLink(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirLink+"/lnk_%d.copy", i)
		_ = KFile.CopyLink(fileLink, des)
	}
}

func TestFile_CopyDir(t *testing.T) {
	var res int64
	var err error

	//忽略已存在的
	res, err = KFile.CopyDir(dirVendor, dirTdat, FILE_COVER_IGNORE)
	assert.Nil(t, err)

	//覆盖已存在的
	res, err = KFile.CopyDir(dirVendor, dirTdat, FILE_COVER_ALLOW)
	assert.Nil(t, err)

	//禁止覆盖
	res, err = KFile.CopyDir(dirVendor, dirTdat, FILE_COVER_DENY)
	assert.Equal(t, int64(0), res)

	//源和目标相同
	res, err = KFile.CopyDir(dirVendor, dirVendor, FILE_COVER_ALLOW)
	assert.Equal(t, int64(0), res)

	//目标为空
	res, err = KFile.CopyDir(dirVendor, "", FILE_COVER_ALLOW)
	assert.NotNil(t, err)

	//源不是目录
	res, err = KFile.CopyDir(fileMd, dirTdat, FILE_COVER_ALLOW)
	assert.NotNil(t, err)
}

func BenchmarkFile_CopyDir(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/copydir_%d", i)
		_, _ = KFile.CopyDir(dirDoc, des, FILE_COVER_ALLOW)
	}
}

func TestFile_DelDir(t *testing.T) {
	var err error
	var chk bool

	//清空目录
	err = KFile.DelDir(dirCopy, false)
	chk = KFile.IsDir(dirCopy)
	assert.Nil(t, err)
	assert.True(t, chk)

	//删除目录
	err = KFile.DelDir(dirNew, true)
	chk = KFile.IsDir(dirNew)
	assert.Nil(t, err)
	assert.False(t, chk)
}

func BenchmarkFile_DelDir(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/copydir_%d", i)
		_ = KFile.DelDir(des, true)
	}
}

func TestFile_Img2Base64(t *testing.T) {
	var res string
	var err error

	//png
	res, err = KFile.Img2Base64(imgPng)
	assert.Nil(t, err)
	assert.Contains(t, res, "png")

	//jpg
	res, err = KFile.Img2Base64(imgJpg)
	assert.Nil(t, err)
	assert.Contains(t, res, "jpg")

	//非图片
	res, err = KFile.Img2Base64(fileMd)
	assert.NotNil(t, err)

	//图片不存在
	res, err = KFile.Img2Base64(fileNone)
	assert.NotNil(t, err)
}

func BenchmarkFile_Img2Base64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Img2Base64(imgPng)
	}
}

func TestFile_FileTree(t *testing.T) {
	var res []string

	//显示全部
	res = KFile.FileTree(dirVendor, FILE_TREE_ALL, true)
	assert.NotEmpty(t, res)

	//仅目录
	res = KFile.FileTree(dirVendor, FILE_TREE_DIR, true)
	assert.NotEmpty(t, res)

	//仅文件
	res = KFile.FileTree(dirVendor, FILE_TREE_FILE, true)
	assert.NotEmpty(t, res)

	//不递归
	res = KFile.FileTree(dirCurr, FILE_TREE_DIR, false)
	assert.GreaterOrEqual(t, len(res), 4)

	//文件过滤
	res = KFile.FileTree(dirCurr, FILE_TREE_FILE, true, func(s string) bool {
		ext := KFile.GetExt(s)
		return ext == "go"
	})
	assert.NotEmpty(t, res)
}

func BenchmarkFile_FileTree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FileTree(dirCurr, FILE_TREE_ALL, false)
	}
}

func TestFile_FormatDir(t *testing.T) {
	var res string

	res = KFile.FormatDir(pathTes3)
	assert.NotContains(t, res, "\\")

	//win格式
	res = KFile.FormatDir(pathTes2)
	assert.Equal(t, 1, strings.Count(res, ":"))

	//空目录
	res = KFile.FormatDir("")
	assert.Empty(t, res)
}

func BenchmarkFile_FormatDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FormatDir(pathTes3)
	}
}

func TestFile_FormatPath(t *testing.T) {
	var res string

	res = KFile.FormatPath(pathTes1)
	assert.NotContains(t, res, ":")

	res = KFile.FormatPath(fileGmod)
	assert.Equal(t, res, fileGmod)

	res = KFile.FormatPath(fileGo)
	assert.Equal(t, res, fileGo)

	res = KFile.FormatPath(pathTes3)
	assert.NotContains(t, res, "\\")

	//win格式
	res = KFile.FormatPath(pathTes2)
	assert.Equal(t, 1, strings.Count(res, ":"))

	//空路径
	res = KFile.FormatPath("")
	assert.Empty(t, res)
}

func BenchmarkFile_FormatPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FormatPath(pathTes2)
	}
}

func TestFile_Md5(t *testing.T) {
	var res string
	var err error

	res, err = KFile.Md5(fileMd, 32)
	assert.NotEmpty(t, res)

	res, err = KFile.Md5(fileMd, 16)
	assert.Nil(t, err)

	//不存在的文件
	res, err = KFile.Md5(fileNone, 32)
	assert.NotNil(t, err)
}

func BenchmarkFile_Md5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Md5(fileMd, 32)
	}
}

func TestFile_ShaX(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res string
	var err error

	res, err = KFile.ShaX(fileGmod, 1)
	assert.NotEmpty(t, res)

	res, err = KFile.ShaX(fileGmod, 256)
	assert.NotEmpty(t, res)

	res, err = KFile.ShaX(fileGmod, 512)
	assert.NotEmpty(t, res)

	//文件不存在
	res, err = KFile.ShaX(fileNone, 512)
	assert.NotNil(t, err)

	//err x
	res, err = KFile.ShaX(fileGmod, 32)
}

func BenchmarkFile_ShaX(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ShaX(fileGmod, 256)
	}
}

func TestFile_Pathinfo(t *testing.T) {
	var res map[string]string

	//所有信息
	res = KFile.Pathinfo(imgPng, -1)
	assert.Equal(t, 4, len(res))

	//仅目录
	res = KFile.Pathinfo(imgPng, 1)

	//仅基础名(文件+扩展)
	res = KFile.Pathinfo(imgPng, 2)

	//仅扩展名
	res = KFile.Pathinfo(imgPng, 4)

	//仅文件名
	res = KFile.Pathinfo(imgPng, 8)

	//目录+基础名
	res = KFile.Pathinfo(imgPng, 3)

	//特殊类型
	res = KFile.Pathinfo(fileGitkee, -1)
	assert.Empty(t, res["filename"])
}

func BenchmarkFile_Pathinfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.Pathinfo(imgPng, -1)
	}
}

func TestFile_Basename(t *testing.T) {
	var res string

	res = KFile.Basename(fileMd)
	assert.Equal(t, "README.md", res)

	res = KFile.Basename(fileNone)
	assert.Equal(t, "none", res)

	res = KFile.Basename("")
	assert.NotEmpty(t, res)
	assert.Equal(t, ".", res)
}

func BenchmarkFile_Basename(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.Basename(fileDante)
	}
}
