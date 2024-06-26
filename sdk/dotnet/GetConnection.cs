// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetConnection
    {
        /// <summary>
        /// Connection DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myConnection = Airbyte.GetConnection.Invoke(new()
        ///     {
        ///         ConnectionId = "...my_connection_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("airbyte:index/getConnection:getConnection", args ?? new GetConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Connection DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myConnection = Airbyte.GetConnection.Invoke(new()
        ///     {
        ///         ConnectionId = "...my_connection_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("airbyte:index/getConnection:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public string ConnectionId { get; set; } = null!;

        public GetConnectionArgs()
        {
        }
        public static new GetConnectionArgs Empty => new GetConnectionArgs();
    }

    public sealed class GetConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public Input<string> ConnectionId { get; set; } = null!;

        public GetConnectionInvokeArgs()
        {
        }
        public static new GetConnectionInvokeArgs Empty => new GetConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        /// <summary>
        /// A list of configured stream options for a connection.
        /// </summary>
        public readonly Outputs.GetConnectionConfigurationsResult Configurations;
        public readonly string ConnectionId;
        /// <summary>
        /// must be one of ["auto", "us", "eu"]
        /// </summary>
        public readonly string DataResidency;
        public readonly string DestinationId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// Define the location where the data will be stored in the destination. must be one of ["source", "destination", "custom_format"]
        /// </summary>
        public readonly string NamespaceDefinition;
        public readonly string NamespaceFormat;
        /// <summary>
        /// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source. must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]
        /// </summary>
        public readonly string NonBreakingSchemaUpdatesBehavior;
        public readonly string Prefix;
        /// <summary>
        /// schedule for when the the connection should run, per the schedule type
        /// </summary>
        public readonly Outputs.GetConnectionScheduleResult Schedule;
        public readonly string SourceId;
        /// <summary>
        /// must be one of ["active", "inactive", "deprecated"]
        /// </summary>
        public readonly string Status;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetConnectionResult(
            Outputs.GetConnectionConfigurationsResult configurations,

            string connectionId,

            string dataResidency,

            string destinationId,

            string id,

            string name,

            string namespaceDefinition,

            string namespaceFormat,

            string nonBreakingSchemaUpdatesBehavior,

            string prefix,

            Outputs.GetConnectionScheduleResult schedule,

            string sourceId,

            string status,

            string workspaceId)
        {
            Configurations = configurations;
            ConnectionId = connectionId;
            DataResidency = dataResidency;
            DestinationId = destinationId;
            Id = id;
            Name = name;
            NamespaceDefinition = namespaceDefinition;
            NamespaceFormat = namespaceFormat;
            NonBreakingSchemaUpdatesBehavior = nonBreakingSchemaUpdatesBehavior;
            Prefix = prefix;
            Schedule = schedule;
            SourceId = sourceId;
            Status = status;
            WorkspaceId = workspaceId;
        }
    }
}
