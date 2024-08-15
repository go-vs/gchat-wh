package messages

import "github.com/go-vs/gchat-wh/types"

type Attachment struct {
	Name              string              `json:"name,omitempty"`
	ContentName       string              `json:"contentName,omitempty"`
	ContentType       string              `json:"contentType,omitempty"`
	ThumbnailUri      string              `json:"thumbnailUri,omitempty"`
	DownloadUri       string              `json:"downloadUri,omitempty"`
	Source            Source              `json:"source,omitempty"`
	AttachmentDataRef *AttachmentDataRef  `json:"attachmentDataRef,omitempty"`
	DriveDataRef      *types.DriveDataRef `json:"driveDataRef,omitempty"`
}

func NewAttachment() *Attachment {
	return &Attachment{}
}

func (a *Attachment) SetName(name string) *Attachment {
	a.Name = name
	return a
}

func (a *Attachment) SetContentName(contentName string) *Attachment {
	a.ContentName = contentName
	return a
}

func (a *Attachment) SetContentType(contentType string) *Attachment {
	a.ContentType = contentType
	return a
}

func (a *Attachment) SetThumbnailUri(thumbnailUri string) *Attachment {
	a.ThumbnailUri = thumbnailUri
	return a
}

func (a *Attachment) SetDownloadUri(downloadUri string) *Attachment {
	a.DownloadUri = downloadUri
	return a
}

func (a *Attachment) SetSource(source Source) *Attachment {
	a.Source = source
	return a
}

func (a *Attachment) SetAttachmentDataRef(attachmentDataRef *AttachmentDataRef) *Attachment {
	a.AttachmentDataRef = attachmentDataRef
	return a
}

func (a *Attachment) SetDriveDataRef(driveDataRef *types.DriveDataRef) *Attachment {
	a.DriveDataRef = driveDataRef
	return a
}

type AttachmentDataRef struct {
	ResourceName          string `json:"resourceName,omitempty"`
	AttachmentUploadToken string `json:"attachmentUploadToken,omitempty"`
}

func NewAttachmentDataRef() *AttachmentDataRef {
	return &AttachmentDataRef{}
}

func (a *AttachmentDataRef) SetResourceName(resourceName string) *AttachmentDataRef {
	a.ResourceName = resourceName
	return a
}

func (a *AttachmentDataRef) SetAttachmentUploadToken(attachmentUploadToken string) *AttachmentDataRef {
	a.AttachmentUploadToken = attachmentUploadToken
	return a
}

type Source string

const (
	SourceDriveFile       Source = "DRIVE_FILE"
	SourceUploadedContent Source = "UPLOADED_CONTENT"
)
