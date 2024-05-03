// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMarketoConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID of your Marketo developer application. See \n\n the docs \n\n for info on how to obtain this.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The Client Secret of your Marketo developer application. See \n\n the docs \n\n for info on how to obtain this.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// Your Marketo Base URL. See \n\n the docs \n\n for info on how to obtain this.
        /// </summary>
        [Input("domainUrl", required: true)]
        public Input<string> DomainUrl { get; set; } = null!;

        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceMarketoConfigurationArgs()
        {
        }
        public static new SourceMarketoConfigurationArgs Empty => new SourceMarketoConfigurationArgs();
    }
}