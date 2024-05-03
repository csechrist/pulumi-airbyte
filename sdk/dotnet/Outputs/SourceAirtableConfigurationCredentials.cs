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
    public sealed class SourceAirtableConfigurationCredentials
    {
        public readonly Outputs.SourceAirtableConfigurationCredentialsOAuth20? OAuth20;
        public readonly Outputs.SourceAirtableConfigurationCredentialsPersonalAccessToken? PersonalAccessToken;

        [OutputConstructor]
        private SourceAirtableConfigurationCredentials(
            Outputs.SourceAirtableConfigurationCredentialsOAuth20? oAuth20,

            Outputs.SourceAirtableConfigurationCredentialsPersonalAccessToken? personalAccessToken)
        {
            OAuth20 = oAuth20;
            PersonalAccessToken = personalAccessToken;
        }
    }
}