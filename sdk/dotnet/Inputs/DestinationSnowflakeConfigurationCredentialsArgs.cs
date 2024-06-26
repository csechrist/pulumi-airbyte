// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationSnowflakeConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("keyPairAuthentication")]
        public Input<Inputs.DestinationSnowflakeConfigurationCredentialsKeyPairAuthenticationArgs>? KeyPairAuthentication { get; set; }

        [Input("oAuth20")]
        public Input<Inputs.DestinationSnowflakeConfigurationCredentialsOAuth20Args>? OAuth20 { get; set; }

        [Input("usernameAndPassword")]
        public Input<Inputs.DestinationSnowflakeConfigurationCredentialsUsernameAndPasswordArgs>? UsernameAndPassword { get; set; }

        public DestinationSnowflakeConfigurationCredentialsArgs()
        {
        }
        public static new DestinationSnowflakeConfigurationCredentialsArgs Empty => new DestinationSnowflakeConfigurationCredentialsArgs();
    }
}
