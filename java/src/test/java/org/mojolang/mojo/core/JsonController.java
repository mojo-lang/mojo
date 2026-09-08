package org.mojolang.mojo.core;

import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Collections;

@RestController
public class JsonController {
    @GetMapping("/simple")
    public org.mojolang.mojo.core.Object simple () {
        return Object.newBuilder()
                .putVals("Hello", ValueUtil.of(" World"))
                .build();
    }

    @GetMapping("/simple/result")
    public Result<org.mojolang.mojo.core.Object> simpleResult () {
        return Result.build(200,
                "OK",
                Object.newBuilder().putVals("Hello", ValueUtil.of(" World")).build());
    }

    @GetMapping("/simple/pagination")
    public Pagination<org.mojolang.mojo.core.Object> simplePagination () {
        return Pagination.build(
                200,
                "OK",
                Collections.singletonList(Object.newBuilder().putVals("Hello", ValueUtil.of(" World")).build()),
                100,
                "1");
    }

    @RequestMapping(
            value = "/request-body",
            method = RequestMethod.POST,
            consumes = MediaType.APPLICATION_JSON_VALUE
    )
    @ResponseBody
    public org.mojolang.mojo.core.Object requestBody (@RequestBody org.mojolang.mojo.core.Object req) {
        return Object.newBuilder()
                .putVals("Hello", req.getValsOrDefault("name", ValueUtil.of("unknown")))
                .build();
    }

    @RequestMapping(
            value = "/request-list-body",
            method = RequestMethod.POST,
            consumes = MediaType.APPLICATION_JSON_VALUE
    )
    @ResponseBody
    public List<org.mojolang.mojo.core.Object> requestListBody (@RequestBody List<org.mojolang.mojo.core.Object> req) {
        List<org.mojolang.mojo.core.Object> res = new ArrayList<>();
        for (org.mojolang.mojo.core.Object obj : req) {
            res.add(Object.newBuilder()
                    .putVals("Hello", obj.getValsOrDefault("name", ValueUtil.of("unknown")))
                    .build());
        }
        return res;
    }

//    @RequestMapping(
//            value = "/mixed-req",
//            method = RequestMethod.POST,
//            consumes = MediaType.APPLICATION_JSON_VALUE
//    )
//    @ResponseBody
//    public NormalRes mixedReq (@RequestBody MixedReq req) {
//        NormalRes res = new NormalRes();
//        res.setMessage("hello " + req.getReq().getName());
//        return res;
//    }
//
//    @RequestMapping(
//            value = "/mixed",
//            method = RequestMethod.POST,
//            consumes = MediaType.APPLICATION_JSON_VALUE
//    )
//    @ResponseBody
//    public MixedRes mixed (@RequestBody MixedReq req) {
//        MixedRes res = new MixedRes();
//        res.setRes(
//                RequestBodyRes.newBuilder()
//                        .setMessage("hello " + req.getReq().getName())
//                        .build()
//        );
//        return res;
//    }
}