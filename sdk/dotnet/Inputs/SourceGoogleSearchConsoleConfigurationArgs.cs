// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleSearchConsoleConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorization", required: true)]
        public Input<Inputs.SourceGoogleSearchConsoleConfigurationAuthorizationArgs> Authorization { get; set; } = null!;

        /// <summary>
        /// (DEPRCATED) A JSON array describing the custom reports you want to sync from Google Search Console. See our \n\ndocumentation\n\n for more information on formulating custom reports.
        /// </summary>
        [Input("customReports")]
        public Input<string>? CustomReports { get; set; }

        [Input("customReportsArrays")]
        private InputList<Inputs.SourceGoogleSearchConsoleConfigurationCustomReportsArrayArgs>? _customReportsArrays;

        /// <summary>
        /// You can add your Custom Analytics report by creating one.
        /// </summary>
        public InputList<Inputs.SourceGoogleSearchConsoleConfigurationCustomReportsArrayArgs> CustomReportsArrays
        {
            get => _customReportsArrays ?? (_customReportsArrays = new InputList<Inputs.SourceGoogleSearchConsoleConfigurationCustomReportsArrayArgs>());
            set => _customReportsArrays = value;
        }

        /// <summary>
        /// If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our \n\nfull documentation\n\n. must be one of ["final", "all"]; Default: "final"
        /// </summary>
        [Input("dataState")]
        public Input<string>? DataState { get; set; }

        /// <summary>
        /// UTC date in the format YYYY-MM-DD. Any data created after this date will not be replicated. Must be greater or equal to the start date field. Leaving this field blank will replicate all data from the start date onward.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        [Input("siteUrls", required: true)]
        private InputList<string>? _siteUrls;

        /// <summary>
        /// The URLs of the website property attached to your GSC account. Learn more about properties \n\nhere\n\n.
        /// </summary>
        public InputList<string> SiteUrls
        {
            get => _siteUrls ?? (_siteUrls = new InputList<string>());
            set => _siteUrls = value;
        }

        /// <summary>
        /// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated. Default: "2021-01-01"
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceGoogleSearchConsoleConfigurationArgs()
        {
        }
        public static new SourceGoogleSearchConsoleConfigurationArgs Empty => new SourceGoogleSearchConsoleConfigurationArgs();
    }
}