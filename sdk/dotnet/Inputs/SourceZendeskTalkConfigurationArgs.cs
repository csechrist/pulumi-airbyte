// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceZendeskTalkConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Zendesk service provides two authentication methods. Choose between: `OAuth2.0` or `API token`.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.SourceZendeskTalkConfigurationCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// The date from which you'd like to replicate data for Zendesk Talk API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        /// <summary>
        /// This is your Zendesk subdomain that can be found in your account URL. For example, in https://{MY*SUBDOMAIN}.zendesk.com/, where MY*SUBDOMAIN is the value of your subdomain.
        /// </summary>
        [Input("subdomain", required: true)]
        public Input<string> Subdomain { get; set; } = null!;

        public SourceZendeskTalkConfigurationArgs()
        {
        }
        public static new SourceZendeskTalkConfigurationArgs Empty => new SourceZendeskTalkConfigurationArgs();
    }
}
