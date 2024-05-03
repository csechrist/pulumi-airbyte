// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationMssqlConfigurationTunnelMethodPasswordAuthenticationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname of the jump server host that allows inbound ssh tunnel.
        /// </summary>
        [Input("tunnelHost", required: true)]
        public Input<string> TunnelHost { get; set; } = null!;

        /// <summary>
        /// Port on the proxy/jump server that accepts inbound ssh connections. Default: 22
        /// </summary>
        [Input("tunnelPort")]
        public Input<int>? TunnelPort { get; set; }

        /// <summary>
        /// OS-level username for logging into the jump server host
        /// </summary>
        [Input("tunnelUser", required: true)]
        public Input<string> TunnelUser { get; set; } = null!;

        [Input("tunnelUserPassword", required: true)]
        private Input<string>? _tunnelUserPassword;

        /// <summary>
        /// OS-level password for logging into the jump server host
        /// </summary>
        public Input<string>? TunnelUserPassword
        {
            get => _tunnelUserPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tunnelUserPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public DestinationMssqlConfigurationTunnelMethodPasswordAuthenticationGetArgs()
        {
        }
        public static new DestinationMssqlConfigurationTunnelMethodPasswordAuthenticationGetArgs Empty => new DestinationMssqlConfigurationTunnelMethodPasswordAuthenticationGetArgs();
    }
}