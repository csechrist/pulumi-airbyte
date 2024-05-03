// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceConfigcat
    {
        /// <summary>
        /// SourceConfigcat DataSource
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
        ///     var mySourceConfigcat = Airbyte.GetSourceConfigcat.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceConfigcatResult> InvokeAsync(GetSourceConfigcatArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceConfigcatResult>("airbyte:index/getSourceConfigcat:getSourceConfigcat", args ?? new GetSourceConfigcatArgs(), options.WithDefaults());

        /// <summary>
        /// SourceConfigcat DataSource
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
        ///     var mySourceConfigcat = Airbyte.GetSourceConfigcat.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceConfigcatResult> Invoke(GetSourceConfigcatInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceConfigcatResult>("airbyte:index/getSourceConfigcat:getSourceConfigcat", args ?? new GetSourceConfigcatInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceConfigcatArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceConfigcatArgs()
        {
        }
        public static new GetSourceConfigcatArgs Empty => new GetSourceConfigcatArgs();
    }

    public sealed class GetSourceConfigcatInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceConfigcatInvokeArgs()
        {
        }
        public static new GetSourceConfigcatInvokeArgs Empty => new GetSourceConfigcatInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceConfigcatResult
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
        private GetSourceConfigcatResult(
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