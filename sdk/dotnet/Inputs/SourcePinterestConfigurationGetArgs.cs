// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePinterestConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        public Input<Inputs.SourcePinterestConfigurationCredentialsGetArgs>? Credentials { get; set; }

        [Input("customReports")]
        private InputList<Inputs.SourcePinterestConfigurationCustomReportGetArgs>? _customReports;

        /// <summary>
        /// A list which contains ad statistics entries, each entry must have a name and can contains fields, breakdowns or action*breakdowns. Click on "add" to fill this field.
        /// </summary>
        public InputList<Inputs.SourcePinterestConfigurationCustomReportGetArgs> CustomReports
        {
            get => _customReports ?? (_customReports = new InputList<Inputs.SourcePinterestConfigurationCustomReportGetArgs>());
            set => _customReports = value;
        }

        /// <summary>
        /// A date in the format YYYY-MM-DD. If you have not set a date, it would be defaulted to latest allowed date by api (89 days from today).
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        [Input("statuses")]
        private InputList<string>? _statuses;

        /// <summary>
        /// For the ads, ad_groups, and campaigns streams, specifying a status will filter out records that do not match the specified ones. If a status is not specified, the source will default to records with a status of either ACTIVE or PAUSED.
        /// </summary>
        public InputList<string> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<string>());
            set => _statuses = value;
        }

        public SourcePinterestConfigurationGetArgs()
        {
        }
        public static new SourcePinterestConfigurationGetArgs Empty => new SourcePinterestConfigurationGetArgs();
    }
}
