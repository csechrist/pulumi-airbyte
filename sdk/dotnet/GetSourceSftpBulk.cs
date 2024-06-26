// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceSftpBulk
    {
        /// <summary>
        /// SourceSftpBulk DataSource
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
        ///     var mySourceSftpbulk = Airbyte.GetSourceSftpBulk.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceSftpBulkResult> InvokeAsync(GetSourceSftpBulkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceSftpBulkResult>("airbyte:index/getSourceSftpBulk:getSourceSftpBulk", args ?? new GetSourceSftpBulkArgs(), options.WithDefaults());

        /// <summary>
        /// SourceSftpBulk DataSource
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
        ///     var mySourceSftpbulk = Airbyte.GetSourceSftpBulk.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceSftpBulkResult> Invoke(GetSourceSftpBulkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceSftpBulkResult>("airbyte:index/getSourceSftpBulk:getSourceSftpBulk", args ?? new GetSourceSftpBulkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceSftpBulkArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceSftpBulkArgs()
        {
        }
        public static new GetSourceSftpBulkArgs Empty => new GetSourceSftpBulkArgs();
    }

    public sealed class GetSourceSftpBulkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceSftpBulkInvokeArgs()
        {
        }
        public static new GetSourceSftpBulkInvokeArgs Empty => new GetSourceSftpBulkInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceSftpBulkResult
    {
        /// <summary>
        /// The values required to configure the source. Parsed as JSON.
        /// </summary>
        public readonly string Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string SourceId;
        public readonly string SourceType;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceSftpBulkResult(
            string configuration,

            string id,

            string name,

            string sourceId,

            string sourceType,

            string workspaceId)
        {
            Configuration = configuration;
            Id = id;
            Name = name;
            SourceId = sourceId;
            SourceType = sourceType;
            WorkspaceId = workspaceId;
        }
    }
}
