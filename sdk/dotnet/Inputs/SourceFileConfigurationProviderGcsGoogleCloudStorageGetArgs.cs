// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceFileConfigurationProviderGcsGoogleCloudStorageGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described \n\nhere\n\n. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.
        /// </summary>
        [Input("serviceAccountJson")]
        public Input<string>? ServiceAccountJson { get; set; }

        public SourceFileConfigurationProviderGcsGoogleCloudStorageGetArgs()
        {
        }
        public static new SourceFileConfigurationProviderGcsGoogleCloudStorageGetArgs Empty => new SourceFileConfigurationProviderGcsGoogleCloudStorageGetArgs();
    }
}