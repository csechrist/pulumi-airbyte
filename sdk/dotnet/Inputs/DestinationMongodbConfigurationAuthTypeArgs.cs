// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationMongodbConfigurationAuthTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Login/Password.
        /// </summary>
        [Input("loginPassword")]
        public Input<Inputs.DestinationMongodbConfigurationAuthTypeLoginPasswordArgs>? LoginPassword { get; set; }

        /// <summary>
        /// None.
        /// </summary>
        [Input("none")]
        public Input<Inputs.DestinationMongodbConfigurationAuthTypeNoneArgs>? None { get; set; }

        public DestinationMongodbConfigurationAuthTypeArgs()
        {
        }
        public static new DestinationMongodbConfigurationAuthTypeArgs Empty => new DestinationMongodbConfigurationAuthTypeArgs();
    }
}
