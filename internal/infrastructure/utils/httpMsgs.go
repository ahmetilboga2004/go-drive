package utils

const (
	InvalidData            = "The provided data is invalid."                                                    // Geçersiz veri
	MissingInformation     = "Required information is missing or incomplete."                                   // Eksik veya hatalı bilgi
	ValidationRequired     = "The data did not pass validation. Please check and try again."                    // Doğrulama gerekli
	AuthenticationRequired = "Please log in to access this resource."                                           // Oturum açmanız gerekli
	AuthorizationRequired  = "You do not have permission to access this resource."                              // Yetkilendirme gerekli
	ActionNotAllowed       = "You are not allowed to perform this action."                                      // Bu işlem yapılamaz
	SessionExpired         = "Your session has expired. Please log in again."                                   // Oturum süresi dolmuş
	ResourceNotFound       = "The requested resource could not be found."                                       // Kaynak bulunamadı
	FileTooLarge           = "The uploaded file exceeds the allowed size limit."                                // Dosya çok büyük
	UnsupportedFileType    = "The uploaded file type is not supported."                                         // Desteklenmeyen dosya tipi
	FileUploadFailed       = "An error occurred while uploading the file. Please try again."                    // Dosya yükleme hatası
	FileNotFound           = "The specified file could not be found."                                           // Dosya bulunamadı
	FileAccessDenied       = "You do not have permission to access this file."                                  // Dosyaya erişim reddedildi
	InvalidFileFormat      = "The file format is invalid or corrupted."                                         // Geçersiz dosya formatı
	FileAlreadyExists      = "A file with the same name already exists. Please rename your file and try again." // Dosya zaten var
	FileDeleteFailed       = "An error occurred while trying to delete the file."                               // Dosya silme hatası
	FileDownloadFailed     = "An error occurred while downloading the file."                                    // Dosya indirme hatası
	RequestTimeout         = "The request took too long to process. Please try again later."                    // İstek zaman aşımına uğradı
	RateLimitExceeded      = "Too many requests. Please slow down and try again."                               // Çok fazla istek
	InternalServerError    = "An unexpected error occurred. Please try again later."                            // Sunucu hatası
	ServiceUnavailable     = "The service is temporarily unavailable. Please try again later."                  // Hizmet geçici olarak kullanılamıyor
)

const (
	FileUploadedSuccessfully     = "File uploaded successfully."              // Dosya başarıyla yüklendi
	FileDownloadedSuccessfully   = "File downloaded successfully."            // Dosya başarıyla indirildi
	FileDeletedSuccessfully      = "File deleted successfully."               // Dosya başarıyla silindi
	UserAuthenticated            = "User authenticated successfully."         // Kullanıcı başarıyla kimlik doğrulandı
	FileUpdatedSuccessfully      = "File updated successfully."               // Dosya başarıyla güncellendi
	RequestProcessedSuccessfully = "Your request was processed successfully." // İstek başarıyla işlendi
	UserRegisteredSuccessfully   = "User registered successfully."            // Kullanıcı başarıyla kaydedildi
	ActionCompletedSuccessfully  = "The action was completed successfully."   // İşlem başarıyla tamamlandı
)
