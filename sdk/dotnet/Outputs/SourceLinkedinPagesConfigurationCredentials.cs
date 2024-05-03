// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceLinkedinPagesConfigurationCredentials
    {
        public readonly Outputs.SourceLinkedinPagesConfigurationCredentialsAccessToken? AccessToken;
        public readonly Outputs.SourceLinkedinPagesConfigurationCredentialsOAuth20? OAuth20;

        [OutputConstructor]
        private SourceLinkedinPagesConfigurationCredentials(
            Outputs.SourceLinkedinPagesConfigurationCredentialsAccessToken? accessToken,

            Outputs.SourceLinkedinPagesConfigurationCredentialsOAuth20? oAuth20)
        {
            AccessToken = accessToken;
            OAuth20 = oAuth20;
        }
    }
}