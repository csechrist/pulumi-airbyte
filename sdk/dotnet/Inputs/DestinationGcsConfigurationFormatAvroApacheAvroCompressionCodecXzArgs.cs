// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecXzArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// must be one of ["xz"]; Default: "xz"
        /// </summary>
        [Input("codec")]
        public Input<string>? Codec { get; set; }

        /// <summary>
        /// The presets 0-3 are fast presets with medium compression. The presets 4-6 are fairly slow presets with high compression. The default preset is 6. The presets 7-9 are like the preset 6 but use bigger dictionaries and have higher compressor and decompressor memory requirements. Unless the uncompressed size of the file exceeds 8 MiB, 16 MiB, or 32 MiB, it is waste of memory to use the presets 7, 8, or 9, respectively. Read more &lt;a href="https://commons.apache.org/proper/commons-compress/apidocs/org/apache/commons/compress/compressors/xz/XZCompressorOutputStream.html#XZCompressorOutputStream-java.io.OutputStream-int-"&gt;here&lt;/a&gt; for details. Default: 6
        /// </summary>
        [Input("compressionLevel")]
        public Input<int>? CompressionLevel { get; set; }

        public DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecXzArgs()
        {
        }
        public static new DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecXzArgs Empty => new DestinationGcsConfigurationFormatAvroApacheAvroCompressionCodecXzArgs();
    }
}