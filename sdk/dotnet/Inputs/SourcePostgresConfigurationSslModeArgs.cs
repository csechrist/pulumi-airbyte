// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourcePostgresConfigurationSslModeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables encryption only when required by the source database.
        /// </summary>
        [Input("allow")]
        public Input<Inputs.SourcePostgresConfigurationSslModeAllowArgs>? Allow { get; set; }

        /// <summary>
        /// Disables encryption of communication between Airbyte and source database.
        /// </summary>
        [Input("disable")]
        public Input<Inputs.SourcePostgresConfigurationSslModeDisableArgs>? Disable { get; set; }

        /// <summary>
        /// Allows unencrypted connection only if the source database does not support encryption.
        /// </summary>
        [Input("prefer")]
        public Input<Inputs.SourcePostgresConfigurationSslModePreferArgs>? Prefer { get; set; }

        /// <summary>
        /// Always require encryption. If the source database server does not support encryption, connection will fail.
        /// </summary>
        [Input("require")]
        public Input<Inputs.SourcePostgresConfigurationSslModeRequireArgs>? Require { get; set; }

        /// <summary>
        /// Always require encryption and verifies that the source database server has a valid SSL certificate.
        /// </summary>
        [Input("verifyCa")]
        public Input<Inputs.SourcePostgresConfigurationSslModeVerifyCaArgs>? VerifyCa { get; set; }

        /// <summary>
        /// This is the most secure mode. Always require encryption and verifies the identity of the source database server.
        /// </summary>
        [Input("verifyFull")]
        public Input<Inputs.SourcePostgresConfigurationSslModeVerifyFullArgs>? VerifyFull { get; set; }

        public SourcePostgresConfigurationSslModeArgs()
        {
        }
        public static new SourcePostgresConfigurationSslModeArgs Empty => new SourcePostgresConfigurationSslModeArgs();
    }
}