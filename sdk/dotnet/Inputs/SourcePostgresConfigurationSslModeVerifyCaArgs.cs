// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePostgresConfigurationSslModeVerifyCaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parsed as JSON.
        /// </summary>
        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        /// <summary>
        /// CA certificate
        /// </summary>
        [Input("caCertificate", required: true)]
        public Input<string> CaCertificate { get; set; } = null!;

        /// <summary>
        /// Client certificate
        /// </summary>
        [Input("clientCertificate")]
        public Input<string>? ClientCertificate { get; set; }

        [Input("clientKey")]
        private Input<string>? _clientKey;

        /// <summary>
        /// Client key
        /// </summary>
        public Input<string>? ClientKey
        {
            get => _clientKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientKeyPassword")]
        private Input<string>? _clientKeyPassword;

        /// <summary>
        /// Password for keystorage. If you do not add it - the password will be generated automatically.
        /// </summary>
        public Input<string>? ClientKeyPassword
        {
            get => _clientKeyPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientKeyPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourcePostgresConfigurationSslModeVerifyCaArgs()
        {
        }
        public static new SourcePostgresConfigurationSslModeVerifyCaArgs Empty => new SourcePostgresConfigurationSslModeVerifyCaArgs();
    }
}