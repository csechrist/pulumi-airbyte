// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class DestinationS3ConfigurationFormatAvroApacheAvro
    {
        /// <summary>
        /// The compression algorithm used to compress data. Default to no compression.
        /// </summary>
        public readonly Outputs.DestinationS3ConfigurationFormatAvroApacheAvroCompressionCodec CompressionCodec;
        /// <summary>
        /// must be one of ["Avro"]; Default: "Avro"
        /// </summary>
        public readonly string? FormatType;

        [OutputConstructor]
        private DestinationS3ConfigurationFormatAvroApacheAvro(
            Outputs.DestinationS3ConfigurationFormatAvroApacheAvroCompressionCodec compressionCodec,

            string? formatType)
        {
            CompressionCodec = compressionCodec;
            FormatType = formatType;
        }
    }
}
