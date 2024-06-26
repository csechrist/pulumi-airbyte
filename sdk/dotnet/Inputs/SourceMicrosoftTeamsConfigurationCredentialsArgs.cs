// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMicrosoftTeamsConfigurationCredentialsArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticateViaMicrosoft")]
        public Input<Inputs.SourceMicrosoftTeamsConfigurationCredentialsAuthenticateViaMicrosoftArgs>? AuthenticateViaMicrosoft { get; set; }

        [Input("authenticateViaMicrosoftOAuth20")]
        public Input<Inputs.SourceMicrosoftTeamsConfigurationCredentialsAuthenticateViaMicrosoftOAuth20Args>? AuthenticateViaMicrosoftOAuth20 { get; set; }

        public SourceMicrosoftTeamsConfigurationCredentialsArgs()
        {
        }
        public static new SourceMicrosoftTeamsConfigurationCredentialsArgs Empty => new SourceMicrosoftTeamsConfigurationCredentialsArgs();
    }
}
