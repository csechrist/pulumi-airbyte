// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGitlabConfigurationCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("oAuth20")]
        public Input<Inputs.SourceGitlabConfigurationCredentialsOAuth20GetArgs>? OAuth20 { get; set; }

        [Input("privateToken")]
        public Input<Inputs.SourceGitlabConfigurationCredentialsPrivateTokenGetArgs>? PrivateToken { get; set; }

        public SourceGitlabConfigurationCredentialsGetArgs()
        {
        }
        public static new SourceGitlabConfigurationCredentialsGetArgs Empty => new SourceGitlabConfigurationCredentialsGetArgs();
    }
}
