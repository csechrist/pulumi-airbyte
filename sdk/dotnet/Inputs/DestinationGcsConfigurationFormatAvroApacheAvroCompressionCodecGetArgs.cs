// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bzip2")]
        public Input<Inputs.DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecBzip2GetArgs>? Bzip2 { get; set; }

        [Input("deflate")]
        public Input<Inputs.DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecDeflateGetArgs>? Deflate { get; set; }

        [Input("noCompression")]
        public Input<Inputs.DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecNoCompressionGetArgs>? NoCompression { get; set; }

        [Input("snappy")]
        public Input<Inputs.DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecSnappyGetArgs>? Snappy { get; set; }

        [Input("xz")]
        public Input<Inputs.DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecXzGetArgs>? Xz { get; set; }

        [Input("zstandard")]
        public Input<Inputs.DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecZstandardGetArgs>? Zstandard { get; set; }

        public DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecGetArgs()
        {
        }
        public static new DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecGetArgs Empty => new DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecGetArgs();
    }
}