// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAsanaConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticateViaAsanaOauth")]
        public Input<Inputs.SourceAsanaConfigurationCredentialsAuthenticateViaAsanaOauthGetArgs>? AuthenticateViaAsanaOauth { get; set; }

        [Input("authenticateWithPersonalAccessToken")]
        public Input<Inputs.SourceAsanaConfigurationCredentialsAuthenticateWithPersonalAccessTokenGetArgs>? AuthenticateWithPersonalAccessToken { get; set; }

        public SourceAsanaConfigurationCredentialsGetArgs()
        {
        }
        public static new SourceAsanaConfigurationCredentialsGetArgs Empty => new SourceAsanaConfigurationCredentialsGetArgs();
    }
}