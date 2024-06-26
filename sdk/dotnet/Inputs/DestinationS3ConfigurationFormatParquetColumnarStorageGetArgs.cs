// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationS3ConfigurationFormatParquetColumnarStorageGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is the size of a row group being buffered in memory. It limits the memory usage when writing. Larger values will improve the IO when reading, but consume more memory when writing. Default: 128 MB. Default: 128
        /// </summary>
        [Input("blockSizeMb")]
        public Input<int>? BlockSizeMb { get; set; }

        /// <summary>
        /// The compression algorithm used to compress data pages. must be one of ["UNCOMPRESSED", "SNAPPY", "GZIP", "LZO", "BROTLI", "LZ4", "ZSTD"]; Default: "UNCOMPRESSED"
        /// </summary>
        [Input("compressionCodec")]
        public Input<string>? CompressionCodec { get; set; }

        /// <summary>
        /// Default: true. Default: true
        /// </summary>
        [Input("dictionaryEncoding")]
        public Input<bool>? DictionaryEncoding { get; set; }

        /// <summary>
        /// There is one dictionary page per column per row group when dictionary encoding is used. The dictionary page size works like the page size but for dictionary. Default: 1024 KB. Default: 1024
        /// </summary>
        [Input("dictionaryPageSizeKb")]
        public Input<int>? DictionaryPageSizeKb { get; set; }

        /// <summary>
        /// must be one of ["Parquet"]; Default: "Parquet"
        /// </summary>
        [Input("formatType")]
        public Input<string>? FormatType { get; set; }

        /// <summary>
        /// Maximum size allowed as padding to align row groups. This is also the minimum size of a row group. Default: 8 MB. Default: 8
        /// </summary>
        [Input("maxPaddingSizeMb")]
        public Input<int>? MaxPaddingSizeMb { get; set; }

        /// <summary>
        /// The page size is for compression. A block is composed of pages. A page is the smallest unit that must be read fully to access a single record. If this value is too small, the compression will deteriorate. Default: 1024 KB. Default: 1024
        /// </summary>
        [Input("pageSizeKb")]
        public Input<int>? PageSizeKb { get; set; }

        public DestinationS3ConfigurationFormatParquetColumnarStorageGetArgs()
        {
        }
        public static new DestinationS3ConfigurationFormatParquetColumnarStorageGetArgs Empty => new DestinationS3ConfigurationFormatParquetColumnarStorageGetArgs();
    }
}
