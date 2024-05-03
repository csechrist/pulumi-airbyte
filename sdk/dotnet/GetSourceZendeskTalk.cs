// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceZendeskTalk
    {
        /// <summary>
        /// SourceZendeskTalk DataSource
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
        ///     var mySourceZendesktalk = Airbyte.GetSourceZendeskTalk.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceZendeskTalkResult> InvokeAsync(GetSourceZendeskTalkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceZendeskTalkResult>("airbyte:index/getSourceZendeskTalk:getSourceZendeskTalk", args ?? new GetSourceZendeskTalkArgs(), options.WithDefaults());

        /// <summary>
        /// SourceZendeskTalk DataSource
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
        ///     var mySourceZendesktalk = Airbyte.GetSourceZendeskTalk.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceZendeskTalkResult> Invoke(GetSourceZendeskTalkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceZendeskTalkResult>("airbyte:index/getSourceZendeskTalk:getSourceZendeskTalk", args ?? new GetSourceZendeskTalkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceZendeskTalkArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceZendeskTalkArgs()
        {
        }
        public static new GetSourceZendeskTalkArgs Empty => new GetSourceZendeskTalkArgs();
    }

    public sealed class GetSourceZendeskTalkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceZendeskTalkInvokeArgs()
        {
        }
        public static new GetSourceZendeskTalkInvokeArgs Empty => new GetSourceZendeskTalkInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceZendeskTalkResult
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
        private GetSourceZendeskTalkResult(
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