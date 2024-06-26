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
    public sealed class SourceBraintreeConfiguration
    {
        /// <summary>
        /// Environment specifies where the data will come from. must be one of ["Development", "Sandbox", "Qa", "Production"]
        /// </summary>
        public readonly string Environment;
        /// <summary>
        /// The unique identifier for your entire gateway account. See the \n\ndocs\n\n for more information on how to obtain this ID.
        /// </summary>
        public readonly string MerchantId;
        /// <summary>
        /// Braintree Private Key. See the \n\ndocs\n\n for more information on how to obtain this key.
        /// </summary>
        public readonly string PrivateKey;
        /// <summary>
        /// Braintree Public Key. See the \n\ndocs\n\n for more information on how to obtain this key.
        /// </summary>
        public readonly string PublicKey;
        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
        /// </summary>
        public readonly string? StartDate;

        [OutputConstructor]
        private SourceBraintreeConfiguration(
            string environment,

            string merchantId,

            string privateKey,

            string publicKey,

            string? startDate)
        {
            Environment = environment;
            MerchantId = merchantId;
            PrivateKey = privateKey;
            PublicKey = publicKey;
            StartDate = startDate;
        }
    }
}
