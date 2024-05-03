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
    public sealed class SourceZendeskTalkConfigurationCredentials
    {
        public readonly Outputs.SourceZendeskTalkConfigurationCredentialsApiToken? ApiToken;
        public readonly Outputs.SourceZendeskTalkConfigurationCredentialsOAuth20? OAuth20;

        [OutputConstructor]
        private SourceZendeskTalkConfigurationCredentials(
            Outputs.SourceZendeskTalkConfigurationCredentialsApiToken? apiToken,

            Outputs.SourceZendeskTalkConfigurationCredentialsOAuth20? oAuth20)
        {
            ApiToken = apiToken;
            OAuth20 = oAuth20;
        }
    }
}