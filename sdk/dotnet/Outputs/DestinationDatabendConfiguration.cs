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
    public sealed class DestinationDatabendConfiguration
    {
        /// <summary>
        /// Name of the database.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// Hostname of the database.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Password associated with the username.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Port of the database. Default: 443
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The default  table was written to. Default: "default"
        /// </summary>
        public readonly string? Table;
        /// <summary>
        /// Username to use to access the database.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private DestinationDatabendConfiguration(
            string database,

            string host,

            string? password,

            int? port,

            string? table,

            string username)
        {
            Database = database;
            Host = host;
            Password = password;
            Port = port;
            Table = table;
            Username = username;
        }
    }
}