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
    public sealed class DestinationMysqlConfigurationTunnelMethodPasswordAuthentication
    {
        /// <summary>
        /// Hostname of the jump server host that allows inbound ssh tunnel.
        /// </summary>
        public readonly string TunnelHost;
        /// <summary>
        /// Port on the proxy/jump server that accepts inbound ssh connections. Default: 22
        /// </summary>
        public readonly int? TunnelPort;
        /// <summary>
        /// OS-level username for logging into the jump server host
        /// </summary>
        public readonly string TunnelUser;
        /// <summary>
        /// OS-level password for logging into the jump server host
        /// </summary>
        public readonly string TunnelUserPassword;

        [OutputConstructor]
        private DestinationMysqlConfigurationTunnelMethodPasswordAuthentication(
            string tunnelHost,

            int? tunnelPort,

            string tunnelUser,

            string tunnelUserPassword)
        {
            TunnelHost = tunnelHost;
            TunnelPort = tunnelPort;
            TunnelUser = tunnelUser;
            TunnelUserPassword = tunnelUserPassword;
        }
    }
}
