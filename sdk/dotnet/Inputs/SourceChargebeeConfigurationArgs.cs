// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceChargebeeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Product Catalog version of your Chargebee site. Instructions on how to find your version you may find \n\nhere\n\n under `API Version` section. If left blank, the product catalog version will be set to 2.0. must be one of ["1.0", "2.0"]; Default: "2.0"
        /// </summary>
        [Input("productCatalog")]
        public Input<string>? ProductCatalog { get; set; }

        /// <summary>
        /// The site prefix for your Chargebee instance.
        /// </summary>
        [Input("site", required: true)]
        public Input<string> Site { get; set; } = null!;

        [Input("siteApiKey", required: true)]
        private Input<string>? _siteApiKey;

        /// <summary>
        /// Chargebee API Key. See the \n\ndocs\n\n for more information on how to obtain this key.
        /// </summary>
        public Input<string>? SiteApiKey
        {
            get => _siteApiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _siteApiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00.000Z. Any data before this date will not be replicated.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public SourceChargebeeConfigurationArgs()
        {
        }
        public static new SourceChargebeeConfigurationArgs Empty => new SourceChargebeeConfigurationArgs();
    }
}