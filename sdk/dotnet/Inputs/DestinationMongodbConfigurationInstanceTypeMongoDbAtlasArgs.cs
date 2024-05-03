// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationMongodbConfigurationInstanceTypeMongoDbAtlasArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL of a cluster to connect to.
        /// </summary>
        [Input("clusterUrl", required: true)]
        public Input<string> ClusterUrl { get; set; } = null!;

        /// <summary>
        /// must be one of ["atlas"]; Default: "atlas"
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        public DestinationMongodbConfigurationInstanceTypeMongoDbAtlasArgs()
        {
        }
        public static new DestinationMongodbConfigurationInstanceTypeMongoDbAtlasArgs Empty => new DestinationMongodbConfigurationInstanceTypeMongoDbAtlasArgs();
    }
}