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
    public sealed class DestinationFireboltConfiguration
    {
        /// <summary>
        /// Firebolt account to login.
        /// </summary>
        public readonly string? Account;
        /// <summary>
        /// The database to connect to.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// Engine name or url to connect to.
        /// </summary>
        public readonly string? Engine;
        /// <summary>
        /// The host name of your Firebolt database.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// Loading method used to select the way data will be uploaded to Firebolt
        /// </summary>
        public readonly Outputs.DestinationFireboltConfigurationLoadingMethod? LoadingMethod;
        /// <summary>
        /// Firebolt password.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Firebolt email address you use to login.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private DestinationFireboltConfiguration(
            string? account,

            string database,

            string? engine,

            string? host,

            Outputs.DestinationFireboltConfigurationLoadingMethod? loadingMethod,

            string password,

            string username)
        {
            Account = account;
            Database = database;
            Engine = engine;
            Host = host;
            LoadingMethod = loadingMethod;
            Password = password;
            Username = username;
        }
    }
}
