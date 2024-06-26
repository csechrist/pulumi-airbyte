// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSmartsheetsConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials", required: true)]
        public Input<Inputs.SourceSmartsheetsConfigurationCredentialsGetArgs> Credentials { get; set; } = null!;

        [Input("metadataFields")]
        private InputList<string>? _metadataFields;

        /// <summary>
        /// A List of available columns which metadata can be pulled from.
        /// </summary>
        public InputList<string> MetadataFields
        {
            get => _metadataFields ?? (_metadataFields = new InputList<string>());
            set => _metadataFields = value;
        }

        /// <summary>
        /// The spreadsheet ID. Find it by opening the spreadsheet then navigating to File &gt; Properties
        /// </summary>
        [Input("spreadsheetId", required: true)]
        public Input<string> SpreadsheetId { get; set; } = null!;

        /// <summary>
        /// Only rows modified after this date/time will be replicated. This should be an ISO 8601 string, for instance: `2000-01-01T13:00:00`. Default: "2020-01-01T00:00:00+00:00"
        /// </summary>
        [Input("startDatetime")]
        public Input<string>? StartDatetime { get; set; }

        public SourceSmartsheetsConfigurationGetArgs()
        {
        }
        public static new SourceSmartsheetsConfigurationGetArgs Empty => new SourceSmartsheetsConfigurationGetArgs();
    }
}
