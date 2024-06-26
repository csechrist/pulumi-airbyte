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
    public sealed class DestinationDuckdbConfiguration
    {
        /// <summary>
        /// Path to the .duckdb file, or the text 'md:' to connect to MotherDuck. The file will be placed inside that local mount. For more information check out our \n\ndocs\n\n
        /// </summary>
        public readonly string DestinationPath;
        /// <summary>
        /// API key to use for authentication to a MotherDuck database.
        /// </summary>
        public readonly string? MotherduckApiKey;
        /// <summary>
        /// Database schema name, default for duckdb is 'main'.
        /// </summary>
        public readonly string? Schema;

        [OutputConstructor]
        private DestinationDuckdbConfiguration(
            string destinationPath,

            string? motherduckApiKey,

            string? schema)
        {
            DestinationPath = destinationPath;
            MotherduckApiKey = motherduckApiKey;
            Schema = schema;
        }
    }
}
