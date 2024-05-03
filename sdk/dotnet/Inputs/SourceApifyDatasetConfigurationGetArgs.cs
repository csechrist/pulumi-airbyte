// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceApifyDatasetConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the dataset you would like to load to Airbyte. In Apify Console, you can view your datasets in the \n\nStorage section under the Datasets tab\n\n after you login. See the \n\nApify Docs\n\n for more information.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// Personal API token of your Apify account. In Apify Console, you can find your API token in the \n\nSettings section under the Integrations tab\n\n after you login. See the \n\nApify Docs\n\n for more information.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceApifyDatasetConfigurationGetArgs()
        {
        }
        public static new SourceApifyDatasetConfigurationGetArgs Empty => new SourceApifyDatasetConfigurationGetArgs();
    }
}