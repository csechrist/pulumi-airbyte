// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceShopifyConfigurationCredentialsApiPasswordGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiPassword", required: true)]
        private Input<string>? _apiPassword;

        /// <summary>
        /// The API Password for your private application in the `Shopify` store.
        /// </summary>
        public Input<string>? ApiPassword
        {
            get => _apiPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceShopifyConfigurationCredentialsApiPasswordGetArgs()
        {
        }
        public static new SourceShopifyConfigurationCredentialsApiPasswordGetArgs Empty => new SourceShopifyConfigurationCredentialsApiPasswordGetArgs();
    }
}
