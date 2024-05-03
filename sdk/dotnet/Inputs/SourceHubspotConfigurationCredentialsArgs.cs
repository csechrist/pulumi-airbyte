// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceHubspotConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("oAuth")]
        public Input<Inputs.SourceHubspotConfigurationCredentialsOAuthArgs>? OAuth { get; set; }

        [Input("privateApp")]
        public Input<Inputs.SourceHubspotConfigurationCredentialsPrivateAppArgs>? PrivateApp { get; set; }

        public SourceHubspotConfigurationCredentialsArgs()
        {
        }
        public static new SourceHubspotConfigurationCredentialsArgs Empty => new SourceHubspotConfigurationCredentialsArgs();
    }
}