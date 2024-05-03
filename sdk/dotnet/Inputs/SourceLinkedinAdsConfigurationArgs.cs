// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceLinkedinAdsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountIds")]
        private InputList<int>? _accountIds;

        /// <summary>
        /// Specify the account IDs to pull data from, separated by a space. Leave this field empty if you want to pull the data from all accounts accessible by the authenticated user. See the \n\nLinkedIn docs\n\n to locate these IDs.
        /// </summary>
        public InputList<int> AccountIds
        {
            get => _accountIds ?? (_accountIds = new InputList<int>());
            set => _accountIds = value;
        }

        [Input("adAnalyticsReports")]
        private InputList<Inputs.SourceLinkedinAdsConfigurationAdAnalyticsReportArgs>? _adAnalyticsReports;
        public InputList<Inputs.SourceLinkedinAdsConfigurationAdAnalyticsReportArgs> AdAnalyticsReports
        {
            get => _adAnalyticsReports ?? (_adAnalyticsReports = new InputList<Inputs.SourceLinkedinAdsConfigurationAdAnalyticsReportArgs>());
            set => _adAnalyticsReports = value;
        }

        [Input("credentials")]
        public Input<Inputs.SourceLinkedinAdsConfigurationCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceLinkedinAdsConfigurationArgs()
        {
        }
        public static new SourceLinkedinAdsConfigurationArgs Empty => new SourceLinkedinAdsConfigurationArgs();
    }
}