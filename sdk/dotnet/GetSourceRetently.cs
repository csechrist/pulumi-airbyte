// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceRetently
    {
        /// <summary>
        /// SourceRetently DataSource
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
        ///     var mySourceRetently = Airbyte.GetSourceRetently.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceRetentlyResult> InvokeAsync(GetSourceRetentlyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceRetentlyResult>("airbyte:index/getSourceRetently:getSourceRetently", args ?? new GetSourceRetentlyArgs(), options.WithDefaults());

        /// <summary>
        /// SourceRetently DataSource
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
        ///     var mySourceRetently = Airbyte.GetSourceRetently.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceRetentlyResult> Invoke(GetSourceRetentlyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceRetentlyResult>("airbyte:index/getSourceRetently:getSourceRetently", args ?? new GetSourceRetentlyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceRetentlyArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceRetentlyArgs()
        {
        }
        public static new GetSourceRetentlyArgs Empty => new GetSourceRetentlyArgs();
    }

    public sealed class GetSourceRetentlyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceRetentlyInvokeArgs()
        {
        }
        public static new GetSourceRetentlyInvokeArgs Empty => new GetSourceRetentlyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceRetentlyResult
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
        private GetSourceRetentlyResult(
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