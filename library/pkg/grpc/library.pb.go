// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: library/api/proto/library.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_library_api_proto_library_proto protoreflect.FileDescriptor

var file_library_api_proto_library_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a, 0x19, 0x62, 0x6f, 0x6f, 0x6b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xa6, 0x04, 0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x19,
	0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_library_api_proto_library_proto_goTypes = []interface{}{
	(*GetBookInfoRequest)(nil),  // 0: book.GetBookInfoRequest
	(*CreateBookRequest)(nil),   // 1: book.CreateBookRequest
	(*UpdateBookRequest)(nil),   // 2: book.UpdateBookRequest
	(*DeleteBookRequest)(nil),   // 3: book.DeleteBookRequest
	(*GetMovieRequest)(nil),     // 4: movie.GetMovieRequest
	(*CreateMovieRequest)(nil),  // 5: movie.CreateMovieRequest
	(*UpdateMovieRequest)(nil),  // 6: movie.UpdateMovieRequest
	(*DeleteMovieRequest)(nil),  // 7: movie.DeleteMovieRequest
	(*GetBookInfoResponse)(nil), // 8: book.GetBookInfoResponse
	(*CreateBookResponse)(nil),  // 9: book.CreateBookResponse
	(*UpdateBookResponse)(nil),  // 10: book.UpdateBookResponse
	(*DeleteBookResponse)(nil),  // 11: book.DeleteBookResponse
	(*GetMovieResponse)(nil),    // 12: movie.GetMovieResponse
	(*CreateMovieResponse)(nil), // 13: movie.CreateMovieResponse
	(*UpdateMovieResponse)(nil), // 14: movie.UpdateMovieResponse
	(*DeleteMovieResponse)(nil), // 15: movie.DeleteMovieResponse
}
var file_library_api_proto_library_proto_depIdxs = []int32{
	0,  // 0: library.LibraryService.GetBookInfo:input_type -> book.GetBookInfoRequest
	1,  // 1: library.LibraryService.CreateBook:input_type -> book.CreateBookRequest
	2,  // 2: library.LibraryService.UpdateBook:input_type -> book.UpdateBookRequest
	3,  // 3: library.LibraryService.DeleteBook:input_type -> book.DeleteBookRequest
	4,  // 4: library.LibraryService.GetMovie:input_type -> movie.GetMovieRequest
	5,  // 5: library.LibraryService.CreateMovie:input_type -> movie.CreateMovieRequest
	6,  // 6: library.LibraryService.UpdateMovie:input_type -> movie.UpdateMovieRequest
	7,  // 7: library.LibraryService.DeleteMovie:input_type -> movie.DeleteMovieRequest
	8,  // 8: library.LibraryService.GetBookInfo:output_type -> book.GetBookInfoResponse
	9,  // 9: library.LibraryService.CreateBook:output_type -> book.CreateBookResponse
	10, // 10: library.LibraryService.UpdateBook:output_type -> book.UpdateBookResponse
	11, // 11: library.LibraryService.DeleteBook:output_type -> book.DeleteBookResponse
	12, // 12: library.LibraryService.GetMovie:output_type -> movie.GetMovieResponse
	13, // 13: library.LibraryService.CreateMovie:output_type -> movie.CreateMovieResponse
	14, // 14: library.LibraryService.UpdateMovie:output_type -> movie.UpdateMovieResponse
	15, // 15: library.LibraryService.DeleteMovie:output_type -> movie.DeleteMovieResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_library_api_proto_library_proto_init() }
func file_library_api_proto_library_proto_init() {
	if File_library_api_proto_library_proto != nil {
		return
	}
	file_book_api_proto_book_proto_init()
	file_movie_api_proto_movie_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_library_api_proto_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_library_api_proto_library_proto_goTypes,
		DependencyIndexes: file_library_api_proto_library_proto_depIdxs,
	}.Build()
	File_library_api_proto_library_proto = out.File
	file_library_api_proto_library_proto_rawDesc = nil
	file_library_api_proto_library_proto_goTypes = nil
	file_library_api_proto_library_proto_depIdxs = nil
}
