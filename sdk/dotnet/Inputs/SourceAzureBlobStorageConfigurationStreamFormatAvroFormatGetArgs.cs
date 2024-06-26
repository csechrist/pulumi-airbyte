// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAzureBlobStorageConfigurationStreamFormatAvroFormatGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to convert double fields to strings. This is recommended if you have decimal numbers with a high degree of precision because there can be a loss precision when handling floating point numbers. Default: false
        /// </summary>
        [Input("doubleAsString")]
        public Input<bool>? DoubleAsString { get; set; }

        public SourceAzureBlobStorageConfigurationStreamFormatAvroFormatGetArgs()
        {
        }
        public static new SourceAzureBlobStorageConfigurationStreamFormatAvroFormatGetArgs Empty => new SourceAzureBlobStorageConfigurationStreamFormatAvroFormatGetArgs();
    }
}
